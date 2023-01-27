package gnum

import (
	"reflect"
	"sync/atomic"
)

type enumCache struct {
	enumTypeToConfigMap atomic.Value
}

func newEnumCache() *enumCache {
	enumTypeToConfigMap := atomic.Value{}
	enumTypeToConfigMap.Store(make(map[reflect.Type]*enumMetadata, 0))
	return &enumCache{
		enumTypeToConfigMap: enumTypeToConfigMap,
	}
}

// Get return the according config based on the enum definition type.
func (e *enumCache) Get(enumType reflect.Type) (config_ *enumMetadata, ok bool) {
	config_, ok = e.enumTypeToConfigMap.Load().(map[reflect.Type]*enumMetadata)[enumType]
	return
}

// Set adds the existing enumTypeToConfigMap with the new enumMetadata and enum type
// and sets the new mapping to the enumTypeToConfigMap atomic.Value.
func (e *enumCache) Set(enumType reflect.Type, config_ *enumMetadata) {
	currentEnumTypeToConfigMap := e.enumTypeToConfigMap.Load().(map[reflect.Type]*enumMetadata)
	newEnumTypeToConfigMap := make(map[reflect.Type]*enumMetadata, len(currentEnumTypeToConfigMap)+1)
	for k, v := range currentEnumTypeToConfigMap {
		newEnumTypeToConfigMap[k] = v
	}

	newEnumTypeToConfigMap[enumType] = config_
	e.enumTypeToConfigMap.Store(newEnumTypeToConfigMap)
}
