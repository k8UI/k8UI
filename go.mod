module github.com/Nortsx/k8UI

go 1.13

require (
	fyne.io/fyne/v2 v2.0.0
	github.com/vrischmann/envconfig v1.3.0
	k8s.io/apimachinery v0.17.0
	k8s.io/client-go v0.17.0
)
replace github.com/Nortsx/k8UI => github.com/amnid/k8UI init_project_structure
replace github.com/Nortsx/k8UI/cmd => github.com/amnid/k8UI/cmd init_project_structure
