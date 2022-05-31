# Getting Started

## Get the package
`go get github.com/inkit/inkit-go@v0.1.20`

## Import the Inkit Sdk into your project
```golang
import (
	"github.com/inkit/inkit-go/client"
	"github.com/inkit/inkit-go/render"
)
```

## Initialize your Inkit client
If you don't have an Inkit account yet, click here to create an account and grab your Api key: https://app.inkit.com/
```golang
client := client.NewClient("your_api_key")
```

## Create a render
Define your render using CreateRenderOptions
```golang
options := render.CreateRenderOptions{
			Html:   "<html>hello</html>",
			Width:  9,
			Height: 11,
			Unit:   "in",
}

// you can also convert a file from your local file system
options := render.CreateRenderOptions{
			FileName:   "test.html",
			Width:  9,
			Height: 11,
			Unit:   "in",
}
```
Submit your render to Inkit
```golang
render, err := client.Render.Create(&options)
```

## Retrieve your render content

Retrieve data about the render (including status). Make sure your render status is Completed before retrieving file data from it. You will recieve a 404 if the render is still In Progress
```golang
render , err := client.Render.Get(render.Id)
```

Retrieve Pdf
```golang
pdfData, err := client.Render.GetPdf(render.Id)
// or if you want to save it locally to a file
err := client.Render.GetPdfAndSaveToFile(render.Id, "test.pdf")
```

Retrieve Html
```golang
htmlData, err := client.Render.GetHtml(render.Id)
// or if you want to save it locally to a file
err := client.Render.GetHtmlAndSaveToFile(render.Id, "test.html")
```

## View a list of all renders in your organization

```golang

// You can provide pagination options if you want to overwrite the defaults
/*
options := render.RenderListOptions{
	Page:     1,
	PageSize: 1,
}
*/
renders, err := client.Render.List(nil)

if err != nil {
	fmt.Println(err)
	return
}

for _, x := range renders.Items {
	fmt.Println(x)
}

```
