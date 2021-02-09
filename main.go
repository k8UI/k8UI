package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"os"
)

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}

func main() {
	app := app.New()

	w := app.NewWindow("k8UI")
	content := container.NewMax()
	content.Add(widget.NewButton("Test", func() {}))

	split := container.NewHSplit(makeNav(), content)

	w.SetContent(split)

	//item1 := widget.NewTabItem("Contexts",widget.NewLabel("Context 1"))
	//item2 := widget.NewTabItem("Pods",widget.NewLabel("Hello Fyne222!"))
	//item3 := widget.NewTabItem("Containers",widget.NewLabel("Hello Fyne222!"))
	//item4 := widget.NewTabItem("Port forwards",widget.NewLabel("Hello Fyne222!"))
	//item5 := widget.NewTabItem("Services",widget.NewLabel("Hello Fyne222!"))
	//item6 := widget.NewTabItem("Deployments",widget.NewLabel("Hello Fyne222!"))
	//testLabel := widget.NewLabel("Hello Fyne222!");
	//container := widget.NewTabContainer(item1, item2, item3, item4, item5, item6)
	//w.SetContent(widget.NewVBox(
	//	testLabel,
	//	container,
	//	widget.NewButton("Quit", func() {
	//		app.Quit()
	//	}),
	//	widget.NewButton("Add", func() {
	//		container.Append(widget.NewTabItem("Contexts12312312",widget.NewLabel("Context 13123123")))
	//	}),
	//	widget.NewButton("Remove", func() {
	//		container.RemoveIndex(5)
	//	}),
	//	widget.NewButton("Mouse", func() {
	//		var kubeconfig *string
	//		if home := homeDir(); home != "" {
	//			kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	//		} else {
	//			kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	//		}
	//		flag.Parse()
	//
	//		// use the current context in kubeconfig
	//		config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//		if err != nil {
	//			panic(err.Error())
	//		}
	//
	//		// create the clientset
	//		clientset, err := kubernetes.NewForConfig(config)
	//		if err != nil {
	//			panic(err.Error())
	//		}
	//		for {
	//			pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	//			if err != nil {
	//				panic(err.Error())
	//			}
	//			fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
	//
	//			// Examples for error handling:
	//			// - Use helper functions like e.g. errors.IsNotFound()
	//			// - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
	//			namespace := "default"
	//			pod := "example-xxxxx"
	//			_, err = clientset.CoreV1().Pods(namespace).Get(pod, metav1.GetOptions{})
	//			if errors.IsNotFound(err) {
	//				fmt.Printf("Pod %s in namespace %s not found\n", pod, namespace)
	//			} else if statusError, isStatus := err.(*errors.StatusError); isStatus {
	//				fmt.Printf("Error getting pod %s in namespace %s: %v\n",
	//					pod, namespace, statusError.ErrStatus.Message)
	//			} else if err != nil {
	//				panic(err.Error())
	//			} else {
	//				fmt.Printf("Found pod %s in namespace %s\n", pod, namespace)
	//			}
	//
	//			//time.Sleep(10 * time.Second)
	//		}
	//	}),
	//))
	//
	w.Show()
	app.Run()
}

func makeNav() fyne.CanvasObject {
	tree := &widget.Tree{}

	return container.NewBorder(nil, nil, nil, nil, tree)
}
