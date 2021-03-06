package shell

import (
	"yunion.io/x/onecloud/pkg/mcclient"
	"yunion.io/x/onecloud/pkg/mcclient/modules"
	"yunion.io/x/onecloud/pkg/mcclient/options"
)

func init() {
	R(&options.MetadataListOptions{}, "metadata-list", "List metadatas", func(s *mcclient.ClientSession, opts *options.MetadataListOptions) error {
		params, err := options.ListStructToParams(opts)
		if err != nil {
			return err
		}
		result, err := modules.Metadatas.List(s, params)
		if err != nil {
			return err
		}
		printList(result, []string{})
		return nil
	})
}
