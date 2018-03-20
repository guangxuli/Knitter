/*
Copyright 2018 ZTE Corporation. All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package knitteragtobj

import (
	"github.com/ZTE/Knitter/knitter-agent/domain/role/knitter-agent-role"
)

type KnitterAgtObj struct {
	NetConfigFileRole   knitteragentrole.NetConfigFileRole
	DeployPodRole       knitteragentrole.DeployPodRole
	VlanIDAllocatorRole knitteragentrole.VlanIDAllocatorRole
	PortObjRole         knitteragentrole.PortObjRole
	DetachTriggerRole   knitteragentrole.DetachTriggerRole
}

var knitterAgtObj *KnitterAgtObj

func GetKnitterAgtObjSingleton() *KnitterAgtObj {
	if knitterAgtObj == nil {
		knitterAgtObj = new(KnitterAgtObj)
	}
	return knitterAgtObj
}
