package shell

import (
	"strings"

	"yunion.io/x/jsonutils"

	"yunion.io/x/onecloud/pkg/mcclient/modules"
	"yunion.io/x/onecloud/pkg/util/excelutils"
	"yunion.io/x/onecloud/pkg/util/printutils"
)

func printList(list *modules.ListResult, columns []string) {
	printutils.PrintJSONList(list, columns)
}

func printObject(obj jsonutils.JSONObject) {
	printutils.PrintJSONObject(obj)
}

func printObjectRecursive(obj jsonutils.JSONObject) {
	printutils.PrintJSONObjectRecursive(obj)
}

func printObjectRecursiveEx(obj jsonutils.JSONObject, cb printutils.PrintJSONObjectRecursiveExFunc) {
	printutils.PrintJSONObjectRecursiveEx(obj, cb)
}

func printBatchResults(results []modules.SubmitResult, columns []string) {
	printutils.PrintJSONBatchResults(results, columns)
}

func exportList(list *modules.ListResult, file string, exportKeys string, exportTexts string, columns []string) {
	var keys []string
	var texts []string
	if len(exportKeys) > 0 {
		keys = strings.Split(exportKeys, ",")
		texts = strings.Split(exportTexts, ",")
	} else {
		keys = columns
		texts = columns
	}
	excelutils.ExportFile(list.Data, keys, texts, file)
}
