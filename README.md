# Inkit Go Sdk

## Getting Started

1. Import the Inkit Sdk into your project
```golang
import "github.com/inkitio/gosdk/client"
```

2. Initialize your Inkit client
```golang
client := client.NewClient("your_api_key")
```

3. Create a render
i. Define your render using CreateRenderOptions
`
options := render.CreateRenderOptions{
			Html:   "<html>hello</html>",
			Width:  9,
			Height: 11,
			Unit:   "in",
}
`
ii. Submit your render to Inkit
```golang
render, err := client.Render.Create(&options)
```

4. Retrieve you render content

i. Retrieve Pdf
```golang
pdfData, err := client.Render.GetPdf(render.id)
```

ii. Retrieve Html
```golang
htmlData, err := client.Render.GetHtml(render.id)
```

iii. Retrieve data about the render (including status)
```golang
render , err := client.Render.Get(render.id)
```