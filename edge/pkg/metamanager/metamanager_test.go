/*
Copyright 2018 The KubeEdge Authors.

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

package metamanager

import (
	"testing"

	"github.com/kubeedge/beehive/pkg/common"
	"github.com/kubeedge/beehive/pkg/core"
	beehiveContext "github.com/kubeedge/beehive/pkg/core/context"
	commodule "github.com/kubeedge/kubeedge/edge/pkg/common/modules"
)

// metaModule is metamanager implementation of Module interface
var metaModule core.Module

func init() {
	beehiveContext.InitContext([]string{common.MsgCtxTypeChannel})
	add := &common.ModuleInfo{
		ModuleName: MetaManagerModuleName,
		ModuleType: common.MsgCtxTypeChannel,
	}
	beehiveContext.AddModule(add)
}

func TestNameAndGroup(t *testing.T) {
	modules := core.GetModules()
	core.Register(&metaManager{enable: true})
	for name, module := range modules {
		if name == MetaManagerModuleName {
			metaModule = module.GetModule()
			break
		}
	}
	t.Run("TestNameAndGroup", func(t *testing.T) {
		if metaModule == nil {
			t.Errorf("failed to register to beehive")
			return
		}
		if MetaManagerModuleName != metaModule.Name() {
			t.Errorf("Name of module is not correct wanted: %v and got: %v", MetaManagerModuleName, metaModule.Name())
			return
		}
		if commodule.MetaGroup != metaModule.Group() {
			t.Errorf("Group of module is not correct wanted: %v and got: %v", commodule.MetaGroup, metaModule.Group())
		}
	})
}
