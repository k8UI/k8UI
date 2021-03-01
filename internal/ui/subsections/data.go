package subsections

import "github.com/k8UI/k8UI/internal/kubernetes"

var SubsIndex = map[string][]string{
	"": {"Context1", "Context2", "Context3"},
	// "collections": {"list", "table", "tree"},
	// "containers":  {"apptabs", "border", "box", "center", "grid", "split", "scroll"},
	// "widgets":     {"accordion", "button", "card", "entry", "form", "input", "text", "toolbar", "progress"},
}

func GetSub(uid string) []string {
	contexts, err := kubernetes.GetContexts()
	if err != nil {
		print("Error is " + err.Error())
	}
	return contexts
}

func AddSub() {
	SubsIndex[""] = append(SubsIndex[""], "Context 4")
}
