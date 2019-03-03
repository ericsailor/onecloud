package k8s

import (
	"fmt"

	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/mcclient/modules/k8s"
	o "yunion.io/x/onecloud/pkg/mcclient/options/k8s"
)

func initKubeMachine() {
	cmdN := func(action string) string {
		return fmt.Sprintf("kubemachine-%s", action)
	}
	R(&o.NodeListOptions{}, cmdN("list"), "List k8s node machines", func(s *mcclient.ClientSession, args *o.NodeListOptions) error {
		params, err := args.Params()
		if err != nil {
			return err
		}
		result, err := k8s.KubeMachines.List(s, params)
		if err != nil {
			return err
		}
		printList(result, k8s.KubeMachines.GetColumns(s))
		return nil
	})

	R(&o.MachineCreateOptions{}, cmdN("create"), "Create k8s machine", func(s *mcclient.ClientSession, args *o.MachineCreateOptions) error {
		params := args.Params()
		node, err := k8s.KubeMachines.Create(s, params)
		if err != nil {
			return err
		}
		printObject(node)
		return nil
	})

	R(&o.IdentOptions{}, cmdN("show"), "Show details of a machine", func(s *mcclient.ClientSession, args *o.IdentOptions) error {
		result, err := k8s.KubeMachines.Get(s, args.ID, nil)
		if err != nil {
			return err
		}
		printObject(result)
		return nil
	})

	R(&o.IdentsOptions{}, cmdN("delete"), "Delete machine", func(s *mcclient.ClientSession, args *o.IdentsOptions) error {
		ret := k8s.KubeMachines.BatchDelete(s, args.ID, nil)
		printBatchResults(ret, k8s.KubeMachines.GetColumns(s))
		return nil
	})

	R(&o.IdentOptions{}, cmdN("recreate"), "Re-Create machine when create fail", func(s *mcclient.ClientSession, args *o.IdentOptions) error {
		ret, err := k8s.KubeMachines.PerformAction(s, args.ID, "recreate", nil)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})

	R(&o.IdentOptions{}, cmdN("terminate"), "Terminate a machine", func(s *mcclient.ClientSession, args *o.IdentOptions) error {
		ret, err := k8s.KubeMachines.PerformAction(s, args.ID, "terminate", nil)
		if err != nil {
			return err
		}
		printObject(ret)
		return nil
	})
}
