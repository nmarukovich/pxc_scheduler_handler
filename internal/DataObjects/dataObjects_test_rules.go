package DataObjects

import (
	"pxc_scheduler_handler/internal/Global"
	"reflect"

)


/*
This section contains all the rules for the tests by tested method
 */


/*
evaluateNode section [start]
------------------------
Here we will test all the methods the checker use to decide if a node should go down or up
*/


//WsrepDesync
func rulesTestCheckWsrepDesync (myArgs args, clusterNode testClusterNodeImpl )  []rule {
	testDataNode := myArgs.node

	myRules := []rule{
		{"No change W", clusterNode, args{testDataNode, testDataNode.Hostgroups[0]}, false},
		{"Change in W but only 1", clusterNode, args{
			changeDataObjectIntAttribute(testDataNode, "WsrepStatus", 2),
			changeHostgGroupIntAttribute(testDataNode.Hostgroups[0], "Size", 1)}, false},
		{"Change in W and 2 in HG", clusterNode, args{
			changeDataObjectIntAttribute(testDataNode, "WsrepStatus", 2),
			changeHostgGroupIntAttribute(testDataNode.Hostgroups[0], "Size", 2)}, true},
		{"Change in R and 1 in HG", clusterNode, args{
			changeDataObjectIntAttribute(testDataNode, "WsrepStatus", 2),
			changeHostgGroupStringAttribute(changeHostgGroupIntAttribute(testDataNode.Hostgroups[0], "Size", 1), "Type", "R")}, false},

		{"Change in R and 2 in HG", clusterNode, args{
			changeDataObjectIntAttribute(testDataNode, "WsrepStatus", 2),
			changeHostgGroupStringAttribute(changeHostgGroupIntAttribute(testDataNode.Hostgroups[0], "Size", 2), "Type", "R")}, true},
	}
	return myRules
}

func rulesTestcheckAnyNotReadyStatus (myArgs args, clusterNode testClusterNodeImpl )  []rule {
	testDataNode := myArgs.node


	myRules := []rule{
		{ "No change W", clusterNode,args{testDataNode,testDataNode.Hostgroups[0]},false},
		{ "Change wsrepstatus = 1 in W but only 1", clusterNode,args{
			changeDataObjectIntAttribute(testDataNode,"WsrepStatus",1),
			changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",1)},true},
		{ "Change wsrepstatus = 1 in W and 2 in HG", clusterNode,args{
			changeDataObjectIntAttribute(testDataNode,"WsrepStatus",1),
			changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",2)},true},
		{ "Change wsrepstatus = 1 in R and 1 in HG", clusterNode,args{
			changeDataObjectIntAttribute(testDataNode,"WsrepStatus",1),
			changeHostgGroupStringAttribute(changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",1),"Type","R")},true},

		{ "Change wsrepstatus = 1 in R and 2 in HG", clusterNode,args{
			changeDataObjectIntAttribute(testDataNode,"WsrepStatus",1),
			changeHostgGroupStringAttribute(changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",2),"Type","R")},true},



		{ "Change wsrepstatus = 3 in W but only 1", clusterNode,args{
			changeDataObjectIntAttribute(testDataNode,"WsrepStatus",3),
			changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",1)},true},
		{ "Change wsrepstatus = 3 in W and 2 in HG", clusterNode,args{
			changeDataObjectIntAttribute(testDataNode,"WsrepStatus",3),
			changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",2)},true},
		{ "Change wsrepstatus = 3 in R and 1 in HG", clusterNode,args{
			changeDataObjectIntAttribute(testDataNode,"WsrepStatus",3),
			changeHostgGroupStringAttribute(changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",1),"Type","R")},true},

		{ "Change wsrepstatus = 3 in R and 2 in HG", clusterNode,args{
			changeDataObjectIntAttribute(testDataNode,"WsrepStatus",3),
			changeHostgGroupStringAttribute(changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",2),"Type","R")},true},

	}
	return myRules
}

func rulesTestcheckNotPrimary(myArgs args, clusterNode testClusterNodeImpl) []rule {
		testDataNode := myArgs.node


		myRules := []rule{
			{ "No WsrepClusterStatus change ", clusterNode,args{testDataNode,testDataNode.Hostgroups[0]},false},
			{ "Change WsrepClusterStatus = NOT Primary in W ", clusterNode,args{
				changeDataObjectStringAttribute(testDataNode,"WsrepClusterStatus","TEST"),
				changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",1)},true},
		}
		return myRules
}

func rulesTestcheckRejectQueries (myArgs args, clusterNode testClusterNodeImpl )  []rule {
	testDataNode := myArgs.node


	myRules := []rule{
		{ "No WsrepRejectqueries change ", clusterNode,args{testDataNode,testDataNode.Hostgroups[0]},false},
		{ "Change WsrepRejectqueries = ALL in W ", clusterNode,args{
			changeDataObjectBoolAttribute(testDataNode,"WsrepRejectqueries",true),
			changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",1)},true},
		{ "Change WsrepRejectqueries = ALL in W ", clusterNode,args{
			changeDataObjectBoolAttribute(testDataNode,"WsrepRejectqueries",true),
			changeHostgGroupStringAttribute(testDataNode.Hostgroups[0],"Type","R")},true},

	}
	return myRules
}

func rulesTestcheckDonorReject (myArgs args, clusterNode testClusterNodeImpl )  []rule {
	testDataNode := myArgs.node
	testDataNode = changeDataObjectIntAttribute(testDataNode,"WsrepStatus",2)

	myRules := []rule{
		{ "No WsrepDonorrejectqueries change ", clusterNode,args{testDataNode,testDataNode.Hostgroups[0]},false},
		{ "Change WsrepDonorrejectqueries = true in W HG size = 1", clusterNode,args{
			changeDataObjectBoolAttribute(testDataNode,"WsrepDonorrejectqueries",true),
			changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",1)},false},
		{ "Change WsrepDonorrejectqueries = true in W HG size = 2", clusterNode,args{
			changeDataObjectBoolAttribute(testDataNode,"WsrepDonorrejectqueries",true),
			changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",2)},true},

		{ "Change WsrepDonorrejectqueries = true in R and 1 in HG", clusterNode,args{
			changeDataObjectBoolAttribute(testDataNode,"WsrepDonorrejectqueries",true),
			changeHostgGroupStringAttribute(changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",1),"Type","R")},false},
		{ "Change WsrepDonorrejectqueries = true in R and 2 in HG", clusterNode,args{
			changeDataObjectBoolAttribute(testDataNode,"WsrepDonorrejectqueries",true),
			changeHostgGroupStringAttribute(changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",2),"Type","R")},true},
	}
	return myRules
}

func rulesTestcheckPxcMaint (myArgs args, clusterNode testClusterNodeImpl )  []rule {
	testDataNode := myArgs.node
	testDataNode = changeDataObjectIntAttribute(testDataNode,"WsrepStatus",4)

	myRules := []rule{
		{ "No PxcMaintMode change ", clusterNode,args{testDataNode,testDataNode.Hostgroups[0]},false},
		{ "Change PxcMaintMode = Maintenance", clusterNode,args{
			changeDataObjectStringAttribute(testDataNode,"PxcMaintMode","Maintenance"),
			changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",1)},true},
		{ "Change PxcMaintMode = Maintenance and OFFLINE_SOFT", clusterNode,args{
			changeDataObjectStringAttribute(changeDataObjectStringAttribute(testDataNode,"ProxyStatus","OFFLINE_SOFT"),
				"PxcMaintMode","Maintenance"),
			changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",1)},false},
	}
	return myRules
}

func rulesTestCheckReadOnly (myArgs args, clusterNode testClusterNodeImpl )  []rule {
	testDataNode := myArgs.node


	myRules := []rule{
		{ "No ReadOnly change ", clusterNode,args{testDataNode,testDataNode.Hostgroups[0]},false},
		{ "Change ReadOnly = On 1 writer ", clusterNode,args{
			changeDataObjectBoolAttribute(testDataNode,"ReadOnly",true),
			changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",1)},true},
		{ "Change ReadOnly = On 2 writers ", clusterNode,args{
			changeDataObjectBoolAttribute(testDataNode,"ReadOnly",true),
			changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",2)},true},

		{ "CChange ReadOnly = On 2 readers ", clusterNode,args{
			changeDataObjectBoolAttribute(testDataNode,"ReadOnly",true),
			changeHostgGroupStringAttribute(changeHostgGroupIntAttribute(testDataNode.Hostgroups[0],"Size",2),"Type","R")},true},

	}
	return myRules
}
/*
evaluateNode section [end]
------------------------
 */


func rulesTestCheckBackOffLine (myArgs args, clusterNode testClusterNodeImpl )  []rule {
	testDataNode := myArgs.node
	testDataNode = changeDataObjectIntAttribute(testDataNode,"WsrepStatus",4)
	testDataNode = changeDataObjectStringAttribute(testDataNode,"ProxyStatus","OFFLINE_SOFT")
	testDataNode = changeDataObjectStringAttribute(testDataNode,"WsrepClusterStatus","Primary")
	testDataNode = changeDataObjectStringAttribute(testDataNode,"PxcMaintMode","DISABLED")
	testDataNode = changeDataObjectBoolAttribute(testDataNode,"WsrepRejectqueries",false)

	myRules := []rule{

		{ "Back from OFFLINE_SOFT no changes ", clusterNode,args{testDataNode,testDataNode.Hostgroups[0]},true},
		{ "Back from OFFLINE_SOFT change WsrepStatus", clusterNode,args{changeDataObjectIntAttribute(testDataNode,"WsrepStatus",2),testDataNode.Hostgroups[0]},false},
		{ "Back from OFFLINE_SOFT change ProxyStatus", clusterNode,args{changeDataObjectStringAttribute(testDataNode,"ProxyStatus","OFFLINE_HARD"),testDataNode.Hostgroups[0]},false},
		{ "Back from OFFLINE_SOFT change WsrepClusterStatus", clusterNode,args{changeDataObjectStringAttribute(testDataNode,"WsrepClusterStatus","Not Primary"),testDataNode.Hostgroups[0]},false},
		{ "Back from OFFLINE_SOFT change PxcMaintMode", clusterNode,args{changeDataObjectStringAttribute(testDataNode,"PxcMaintMode","MAINTENANCE"),testDataNode.Hostgroups[0]},false},
		{ "Back from OFFLINE_SOFT change PxcMaintMode", clusterNode,args{changeDataObjectBoolAttribute(testDataNode,"WsrepRejectqueries",true),testDataNode.Hostgroups[0]},false},

	}
	return myRules
}

func rulesTestCheckBackNew (myArgs args, clusterNode testClusterNodeImpl )  []rule {
	testDataNode := myArgs.node


	myRules := []rule{
		{ "Back online Node is new no changes ", clusterNode,args{testDataNode,testDataNode.Hostgroups[0]},false},
		{ "Back online Node is new no changes ", clusterNode,args{changeDataObjectBoolAttribute(testDataNode,"NodeIsNew",true),testDataNode.Hostgroups[0]},true},
	}
	return myRules
}



func rulesTestCheckBackPrimary (myArgs args, clusterNode testClusterNodeImpl )  []rule {
	testDataNode := myArgs.node
	testDataNode = changeDataObjectIntAttribute(testDataNode,"WsrepStatus",4)
	//testDataNode = changeDataObjectStringAttribute(testDataNode,"ProxyStatus","OFFLINE_SOFT")
	testDataNode = changeDataObjectStringAttribute(testDataNode,"WsrepClusterStatus","Primary")
	testDataNode = changeDataObjectStringAttribute(testDataNode,"PxcMaintMode","DISABLED")
	testDataNode = changeDataObjectBoolAttribute(testDataNode,"WsrepRejectqueries",false)
	testDataNode = changeDataObjectIntAttribute(testDataNode,"HostgroupId",9100)

	myRules := []rule{
		{ "Back online Node from 9000 HG no changes", clusterNode,args{testDataNode,testDataNode.Hostgroups[0]},true},
		{ "Back online Node from 9000 HG change WsrepStatus", clusterNode,args{changeDataObjectIntAttribute(testDataNode,"WsrepStatus",2),testDataNode.Hostgroups[0]},false},
		{ "Back online Node from 9000 HG change HostgroupId", clusterNode,args{changeDataObjectIntAttribute(testDataNode,"HostgroupId",100),testDataNode.Hostgroups[0]},false},
		{ "Back online Node from 9000 HG change WsrepClusterStatus", clusterNode,args{changeDataObjectStringAttribute(testDataNode,"WsrepClusterStatus","Not Primary"),testDataNode.Hostgroups[0]},false},
		{ "Back online Node from 9000 HG change PxcMaintMode", clusterNode,args{changeDataObjectStringAttribute(testDataNode,"PxcMaintMode","MAINTENANCE"),testDataNode.Hostgroups[0]},false},
		{ "Back online Node from 9000 HG change WsrepRejectqueries", clusterNode,args{changeDataObjectBoolAttribute(testDataNode,"WsrepRejectqueries",true),testDataNode.Hostgroups[0]},false},


	}
	return myRules
}

/*
Test RULES Cluster method evaluateWriters [start]
*/
func rulesTestCleanWriters (myArgs args,clusterNode testClusterNodeImpl )  []rule {
	testDataNode := myArgs.node


	myRules := []rule{
		{ "Check evaluateWriters - cleanWriters no changes", clusterNode,args{testDataNode,testDataNode.Hostgroups[0]},false},
		{ "Check evaluateWriters - cleanWriters One node is offline", clusterNode,args{changeDataObjectStringAttribute(testDataNode,"ProxyStatus","OFFLINE_SOFT"),testDataNode.Hostgroups[0]},true},
	}
	return myRules
}

/*
Test RULES Cluster method evaluateWriters [end]
*/

/*
Test Cluster method evaluateReaders [start]
*/

func rulesTestProcessUpAndDownReaders (myArgs args, clusterNode testClusterNodeImpl )  []rule {
	testDataNode := myArgs.node


	myRules := []rule{
		{ "Check evaluateReaders - ProcessUpAndDownReaders no changes", clusterNode,args{testDataNode,testDataNode.Hostgroups[0]},false},
		{ "Check evaluateReaders - ProcessUpAndDownReaders MOVE_DOWN_HG_CHANGE", clusterNode,args{changeDataObjectIntAttribute(testDataNode,"ActionType",3001),testDataNode.Hostgroups[0]},true},
		{ "Check evaluateReaders - ProcessUpAndDownReaders MOVE_DOWN_OFFLINE", clusterNode,args{changeDataObjectIntAttribute(testDataNode,"ActionType",3010),testDataNode.Hostgroups[0]},true},
		{ "Check evaluateReaders - ProcessUpAndDownReaders MOVE_TO_MAINTENANCE", clusterNode,args{changeDataObjectIntAttribute(testDataNode,"ActionType",3020),testDataNode.Hostgroups[0]},true},
		{ "Check evaluateReaders - ProcessUpAndDownReaders MOVE_SWAP_READER_TO_WRITER", clusterNode,args{changeDataObjectIntAttribute(testDataNode,"ActionType",5001),testDataNode.Hostgroups[0]},true},
		{ "Check evaluateReaders - ProcessUpAndDownReaders MOVE_UP_OFFLINE", clusterNode,args{changeDataObjectIntAttribute(testDataNode,"ActionType",1000),testDataNode.Hostgroups[0]},true},
		{ "Check evaluateReaders - ProcessUpAndDownReaders MOVE_UP_HG_CHANGE", clusterNode,args{changeDataObjectIntAttribute(testDataNode,"ActionType",1010),testDataNode.Hostgroups[0]},true},
		{ "Check evaluateReaders - ProcessUpAndDownReaders INSERT_READ", clusterNode,args{changeDataObjectIntAttribute(testDataNode,"ActionType",4010),testDataNode.Hostgroups[0]},true},
		{ "Check evaluateReaders - ProcessUpAndDownReaders MOVE_SWAP_WRITER_TO_READER", clusterNode,args{changeDataObjectIntAttribute(testDataNode,"ActionType",5101),testDataNode.Hostgroups[0]},true},
	}
	return myRules
}


func rulesTestProcessWriterIsAlsoReader (myArgs args, clusterNode testClusterNodeImpl )  []rule {
	testDataNode := myArgs.node


	myRules := []rule{
		{ "Check evaluateReaders - processWriterIsAlsoReader no changes", clusterNode,args{testDataNode,testDataNode.Hostgroups[0]},false},
		{ "Check evaluateReaders - processWriterIsAlsoReader WriterIsAlso reader = 0 and 2 readers ", changeClusterObjectIntAttribute(clusterNode,"WriterIsReader",0) ,args{testDataNode,testDataNode.Hostgroups[0]},true},
		{ "Check evaluateReaders - processWriterIsAlsoReader WriterIsAlso reader = 0 and 1 readers ",  changeClusterObjectIntAttribute(clusterNode,"WriterIsReader",100), args{testDataNode,testDataNode.Hostgroups[0]},false},

	}
	return myRules
}


/*
Test Cluster method evaluateReaders [end]
*/

//TEMPLATE

//func rulesTest<Method> (myArgs args, clusterNode testClusterNodeImpl )  []rule {
//	testDataNode := myArgs.node

//
//	myRules := []rule{
//	}
//	return myRules
//}




//This section contains the helpers to run the tests
//======================

type rule struct {
	name               string
	testClusterNodeImp testClusterNodeImpl
	args               args
	want               bool
}

type args struct {
	node      DataNodeImpl
	currentHg Hostgroup
}

type testClusterNodeImpl struct {
	ActiveFailover    int
	FailBack          bool
	ActionNodes       map[string]DataNodeImpl
	BackupReaders     map[string]DataNodeImpl
	BackupWriters     map[string]DataNodeImpl
	BackupHgReaderId  int
	BakcupHgWriterId  int
	CheckTimeout      int
	ClusterIdentifier int
	ClusterSize       int
	HasPrimary        bool
	ClusterName       string
	Comment           string
	config            Global.Configuration
	Debug             bool
	FailOverNode      DataNodeImpl
	HasFailoverNode   bool
	Haswriter         bool
	HgReaderId        int
	HgWriterId        int
	Hostgroups        map[int]Hostgroup
	MainSegment       int
	MonitorPassword   string
	MonitorUser       string
	Name              string
	NodesPxc          *SyncMap
	NodesPxcMaint     []DataNodeImpl
	MaxNumWriters     int
	OffLineReaders    map[string]DataNodeImpl
	OffLineWriters    map[string]DataNodeImpl
	OffLineHgReaderID int
	OffLineHgWriterId int
	ReaderNodes       map[string]DataNodeImpl
	RequireFailover   bool
	RetryDown         int
	RetryUp           int
	Singlenode        bool
	SinglePrimary     bool
	Size              int
	Ssl               *SslCertificates
	Status            int
	WriterIsReader    int
	WriterNodes       map[string]DataNodeImpl
}


func (clusterNode testClusterNodeImpl) testDataNodeFactory() DataNodeImpl{
	testDataNode := clusterNode.testDataNodeFactoryDns("127.0.0.1:3306")
	return testDataNode
}


func (clusterNode testClusterNodeImpl) testDataNodeFactoryDns(dns string) DataNodeImpl{
	testDataNode := DataNodeImpl{
		ActionType: 0,
		NodeIsNew: false,
		RetryUp: 0,
		RetryDown: 0,
		Comment: "comment",
		Compression: 0,
		Connection: nil,
		ConnUsed: 0,
		Debug: true,
		Dns: dns,
		GtidPort: 33062,
		HostgroupId: 100,
		Hostgroups: []Hostgroup{},
		Ip: "127.0.0.1",
		MaxConnection: 300,
		MaxLatency: 100,
		MaxReplicationLag: 100,
		Name: "testNode",
		NodeTCPDown: false,
		Password: "password",
		Port: 3306,
		Processed: true,
		ProcessStatus: 1,
		ProxyStatus: "ONLINE",
		ReadOnly: false,
		Ssl: new(SslCertificates),
		Status: make(map[string]string),
		UseSsl: false,
		User: "user",
		Variables: make(map[string]string),
		Weight: 1,

		//pxc
		PxcMaintMode: "DISABLED",
		WsrepConnected: true,
		WsrepDesinccount: 0,
		WsrepDonorrejectqueries: false,
		WsrepGcommUuid: "",
		WsrepLocalIndex :        1,
		WsrepPcWeight :          1,
		WsrepProvider:           make(map[string]string),
		WsrepReady :             true,
		WsrepRejectqueries:      false,
		WsrepSegment :           1,
		WsrepStatus:             4,
		WsrepClusterSize:        3,
		WsrepClusterName :       "test",
		WsrepClusterStatus:      "Primary",
		WsrepNodeName  :         "testnode",
		HasPrimaryState :        true,
		PxcView          :       PxcClusterView{},
	}
	return testDataNode
}

func testClusterFactory() testClusterNodeImpl {
	cluster := testClusterNodeImpl{
		ActiveFailover:    0,
		FailBack:          false,
		ActionNodes:       make(map[string]DataNodeImpl),
		BackupReaders:     make(map[string]DataNodeImpl),
		BackupWriters:     make(map[string]DataNodeImpl),
		BackupHgReaderId:  8101,
		BakcupHgWriterId:  8100,
		CheckTimeout:      60,
		ClusterIdentifier: 10,
		ClusterSize:       3,
		HasPrimary:        false,
		ClusterName:       "test",
		Comment:           "comment",
		config:            Global.Configuration{},
		Debug:             true,
		FailOverNode:      DataNodeImpl{},
		HasFailoverNode:   false,
		Haswriter:         true,
		HgReaderId:        101,
		HgWriterId:        100,
		Hostgroups:        make(map[int]Hostgroup),
		MainSegment:       1,
		MonitorPassword:   "password",
		MonitorUser:       "monitor",
		Name:              "test",
		NodesPxc:          new(SyncMap) ,
		NodesPxcMaint:     []DataNodeImpl{},
		MaxNumWriters:     1,
		OffLineReaders:    make(map[string]DataNodeImpl),
		OffLineWriters:    make(map[string]DataNodeImpl),
		OffLineHgReaderID: 9101,
		OffLineHgWriterId: 9100,
		ReaderNodes:       make(map[string]DataNodeImpl),
		RequireFailover:   false,
		RetryDown:         0,
		RetryUp:           0,
		Singlenode:        true,
		SinglePrimary:     true,
		Size:              3,
		Ssl:               new(SslCertificates),
		Status:            4,
		WriterIsReader:    1,
		WriterNodes:       make(map[string]DataNodeImpl),
	}
	return cluster
}

func getHg() Hostgroup{
	currentHg := getHgOpt(100,1,"W")
	return currentHg
}

func getHgOpt(id int, size int, hgType string) Hostgroup{
	currentHg :=  Hostgroup{
		Id:    id,
		Size:  size,
		Type:  hgType,
		Nodes: []DataNodeImpl{},
	}

	return currentHg
}


func changeDataObjectIntAttribute(node DataNodeImpl, name string,value int64) DataNodeImpl{
	reflect.ValueOf(&node).Elem().FieldByName(name).SetInt(value)
	return node

}


func changeHostgGroupIntAttribute(hostG Hostgroup, name string,value int64) Hostgroup{
	reflect.ValueOf(&hostG).Elem().FieldByName(name).SetInt(value)
	return hostG

}

func changeHostgGroupStringAttribute(hostG Hostgroup, name string,value string) Hostgroup{
	reflect.ValueOf(&hostG).Elem().FieldByName(name).SetString(value)
	return hostG

}

func changeDataObjectStringAttribute(node DataNodeImpl, name string, value string) DataNodeImpl {
	reflect.ValueOf(&node).Elem().FieldByName(name).SetString(value)
	return node

}

func changeDataObjectBoolAttribute(node DataNodeImpl, name string, value bool) DataNodeImpl {
	reflect.ValueOf(&node).Elem().FieldByName(name).SetBool(value)
	return node

}

func changeClusterObjectIntAttribute(node testClusterNodeImpl, name string,value int64) testClusterNodeImpl{
	reflect.ValueOf(&node).Elem().FieldByName(name).SetInt(value)
	return node

}
func deleteFromMap (cluster testClusterNodeImpl, name string, key string) testClusterNodeImpl{
	delete(cluster.ReaderNodes,key)
	return cluster
}

func (tt *testClusterNodeImpl) clusterNodeImplFactory() *DataClusterImpl {
	cluster := &DataClusterImpl{
		ActiveFailover:    tt.ActiveFailover,
		FailBack:          tt.FailBack,
		ActionNodes:       tt.ActionNodes,
		BackupReaders:     tt.BackupReaders,
		BackupWriters:     tt.BackupWriters,
		BackupHgReaderId:  tt.BackupHgReaderId,
		BakcupHgWriterId:  tt.BakcupHgWriterId,
		CheckTimeout:      tt.CheckTimeout,
		ClusterIdentifier: tt.ClusterIdentifier,
		ClusterSize:       tt.ClusterSize,
		HasPrimary:        tt.HasPrimary,
		ClusterName:       tt.ClusterName,
		Comment:           tt.Comment,
		config:            tt.config,
		Debug:             tt.Debug,
		FailOverNode:      tt.FailOverNode,
		HasFailoverNode:   tt.HasFailoverNode,
		Haswriter:         tt.Haswriter,
		HgReaderId:        tt.HgReaderId,
		HgWriterId:        tt.HgWriterId,
		Hostgroups:        tt.Hostgroups,
		MainSegment:       tt.MainSegment,
		MonitorPassword:   tt.MonitorPassword,
		MonitorUser:       tt.MonitorUser,
		Name:              tt.Name,
		NodesPxc:          tt.NodesPxc,
		NodesPxcMaint:     tt.NodesPxcMaint,
		MaxNumWriters:     tt.MaxNumWriters,
		OffLineReaders:    tt.OffLineReaders,
		OffLineWriters:    tt.OffLineWriters,
		OffLineHgReaderID: tt.OffLineHgReaderID,
		OffLineHgWriterId: tt.OffLineHgWriterId,
		ReaderNodes:       tt.ReaderNodes,
		RequireFailover:   tt.RequireFailover,
		RetryDown:         tt.RetryDown,
		RetryUp:           tt.RetryUp,
		Singlenode:        tt.Singlenode,
		SinglePrimary:     tt.SinglePrimary,
		Size:              tt.Size,
		Ssl:               tt.Ssl,
		Status:            tt.Status,
		WriterIsReader:    tt.WriterIsReader,
		WriterNodes:       tt.WriterNodes,
	}
	return cluster
}