package worker

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/opensvc/oc3/api"
	"github.com/opensvc/oc3/cachekeys"
)

type (
	jobFeedDaemonPing struct {
		*BaseJob

		nodeID     string
		clusterID  string
		callerNode *DBNode

		byObjectID map[string]*DBObject
		byNodeID   map[string]*DBNode

		// clusterNode is a map of cluster nodes (from POST feed daemon ping)
		clusterNode map[string]struct{}

		// clusterObject is a map of object names (from POST feed daemon ping)
		clusterObject map[string]struct{}
	}
)

func newDaemonPing(nodeID string) *jobFeedDaemonPing {
	return &jobFeedDaemonPing{
		BaseJob: &BaseJob{
			name:   "daemonPing",
			detail: "nodeID: " + nodeID,

			cachePendingH:   cachekeys.FeedDaemonPingPendingH,
			cachePendingIDX: nodeID,
		},
		nodeID: nodeID,

		byNodeID:   make(map[string]*DBNode),
		byObjectID: make(map[string]*DBObject),

		clusterNode:   make(map[string]struct{}),
		clusterObject: make(map[string]struct{}),
	}
}

func (d *jobFeedDaemonPing) Operations() []operation {
	return []operation{
		{desc: "daemonPing/dropPending", do: d.dropPending},
		{desc: "daemonPing/getData", do: d.getData},
		{desc: "daemonPing/dbFetchNodes", do: d.dbFetchNodes},
		{desc: "daemonPing/dbFetchObjects", do: d.dbFetchObjects},
		{desc: "daemonPing/dbPingInstances", do: d.dbPingInstances},
		{desc: "daemonPing/dbPingObjects", do: d.dbPingObjects},
		{desc: "daemonPing/cacheObjectsWithoutConfig", do: d.cacheObjectsWithoutConfig},
		{desc: "daemonPing/pushFromTableChanges", do: d.pushFromTableChanges},
	}
}

func (d *jobFeedDaemonPing) getData() error {
	var data api.PostFeedDaemonPing
	if b, err := d.redis.HGet(d.ctx, cachekeys.FeedDaemonPingH, d.nodeID).Bytes(); err != nil {
		return fmt.Errorf("getData: HGET %s %s: %w", cachekeys.FeedDaemonPingH, d.nodeID, err)
	} else if err = json.Unmarshal(b, &data); err != nil {
		return fmt.Errorf("getData: unexpected data from %s %s: %w", cachekeys.FeedDaemonPingH, d.nodeID, err)
	} else {
		for _, nodename := range data.Nodes {
			d.clusterNode[nodename] = struct{}{}
		}

		for _, objectName := range data.Objects {
			d.clusterObject[objectName] = struct{}{}
		}
	}
	return nil
}

// dbFetchNodes fetch nodes (that are associated with caller node ID) from database
// and sets d.byNodeID and d.clusterID.
func (d *jobFeedDaemonPing) dbFetchNodes() (err error) {
	var (
		dbNodes []*DBNode
	)
	if dbNodes, err = d.oDb.clusterNodesFromNodeID(d.ctx, d.nodeID); err != nil {
		return fmt.Errorf("dbFetchNodes %s: %w", d.nodeID, err)
	}
	for _, n := range dbNodes {
		if _, ok := d.clusterNode[n.nodename]; !ok {
			// skipped: not member of posted cluster nodenames
			continue
		}
		d.byNodeID[n.nodeID] = n
	}
	callerNode, ok := d.byNodeID[d.nodeID]
	if !ok {
		return fmt.Errorf("dbFetchNodes source node has been removed")
	}
	d.callerNode = callerNode
	d.clusterID = callerNode.clusterID
	return nil
}

func (d *jobFeedDaemonPing) dbFetchObjects() (err error) {
	var (
		objects []*DBObject
	)
	if objects, err = d.oDb.objectsFromClusterID(d.ctx, d.clusterID); err != nil {
		return fmt.Errorf("dbFetchObjects query node %s (%s) clusterID: %s: %w",
			d.callerNode.nodename, d.nodeID, d.clusterID, err)
	}
	for _, o := range objects {
		if _, ok := d.clusterObject[o.svcname]; !ok {
			// skipped: not member of posted object names
			continue
		}
		d.byObjectID[o.svcID] = o
		slog.Debug(fmt.Sprintf("dbFetchObjects  %s (%s)", o.svcname, o.svcID))
	}
	return nil
}

// dbPingInstances call opensvcDB.instancePingFromNodeID for all db fetched nodes
func (d *jobFeedDaemonPing) dbPingInstances() error {
	for nodeID := range d.byNodeID {
		if ok, err := d.oDb.instancePingFromNodeID(d.ctx, nodeID); err != nil {
			return fmt.Errorf("dbPingInstances: %w", err)
		} else if ok {
			continue
		}
	}
	return nil
}

// dbPingObjects call opensvcDB.objectPing for all db fetched objects
func (d *jobFeedDaemonPing) dbPingObjects() (err error) {
	for objectID, obj := range d.byObjectID {
		objectName := obj.svcname
		if obj.availStatus != "undef" {
			slog.Debug(fmt.Sprintf("ping svc %s %s", objectName, objectID))
			if _, err := d.oDb.objectPing(d.ctx, objectID); err != nil {
				return fmt.Errorf("dbPingObjects can't ping object %s %s: %w", objectName, objectID, err)
			}
		}
	}
	return nil
}

// cacheObjectsWithoutConfig populate FeedObjectConfigForClusterIDH with names of objects without config
func (d *jobFeedDaemonPing) cacheObjectsWithoutConfig() error {
	objects, err := d.populateFeedObjectConfigForClusterIDH(d.clusterID, d.byObjectID)
	if len(objects) > 0 {
		slog.Info(fmt.Sprintf("daemonPing nodeID: %s need object config: %s", d.nodeID, objects))
	}
	return err
}
