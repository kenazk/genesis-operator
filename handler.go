package main

import (
	log "github.com/Sirupsen/logrus"
	 _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

// Handler interface contains the methods that are required
type Handler interface {
	Init() error
	ObjectCreated(obj interface{})
	ObjectDeleted(obj interface{})
	ObjectUpdated(objOld, objNew interface{})
}

// TestHandler is a sample implementation of Handler
type TestHandler struct{}

// Init handles any handler initialization
func (t *TestHandler) Init() error {
	log.Info("TestHandler.Init")
	return nil
}

// ObjectCreated is called when an object is created
func (t *TestHandler) ObjectCreated(obj interface{}) {
	log.Info("TestHandler.ObjectCreated")
	log.Info("[genesis-log] Creating an instance of Genesis AppStack")
	log.Info("[genesis-log] Creating genesis::aws::vpc")
	log.Info("[genesis-log] Creating genesis::aws::securitygroup")
	log.Info("[genesis-log] Creating genesis::aws::launchconfig")
	log.Info("[genesis-log] Creating genesis::aws::autoscalinggroup")
	log.Info("[genesis-log] Created instance of Genesis AppStack")
}

// ObjectDeleted is called when an object is deleted
func (t *TestHandler) ObjectDeleted(obj interface{}) {
	log.Info("TestHandler.ObjectDeleted")
	log.Info("[genesis-log] Deleting resources vpc, securitygroup")
	log.Info("[genesis-log] Deleting resources launchconfig, autoscalinggroup") 
}

// ObjectUpdated is called when an object is updated
func (t *TestHandler) ObjectUpdated(objOld, objNew interface{}) {
	log.Info("TestHandler.ObjectUpdated")
}
