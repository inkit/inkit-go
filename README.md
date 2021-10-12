# Getting Started

## Import the Inkit Sdk into your project
```golang
import "github.com/inkitio/gosdk/client"
```

## Initialize your Inkit client
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
```
Submit your render to Inkit
```golang
render, err := client.Render.Create(&options)
```

## Retrieve you render content

Retrieve Pdf
```golang
pdfData, err := client.Render.GetPdf(render.id)
// or if you want to save it locally to a file
err := client.Render.GetPdfAndSaveToFile(render.id, "test.pdf")
```

Retrieve Html
```golang
htmlData, err := client.Render.GetHtml(render.id)
// or if you want to save it locally to a file
err := client.Render.GetHtmlAndSaveToFile(render.id, "test.html")
```

Retrieve data about the render (including status)
```golang
render , err := client.Render.Get(render.id)
```