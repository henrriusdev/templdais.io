// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package pages

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/henrriusdev/templdais"

var accordionCode = templdais.AccordionAttrs{
	Items: []templdais.AccordionItem{
		{Title: "Preview", Content: PreviewCode(), Open: true},
		{Title: "Code", Content: CodeComponentCode()},
		{Title: "Attributes", Content: AttributesCode()},
	},
	Name:  "code",
	Arrow: true,
}

var tableCodeAttrs = templdais.TableAttrs{
	TableName: "templdais.CodeAttrs",
	Columns:   []string{"Attribute", "Type", "Description", "Default"},
	Rows: []map[string]string{
		{"Attribute": "Prefix", "Type": "string", "Description": "The prefix of the code", "Default": "empty"},
		{"Attribute": "Brand", "Type": "string", "Description": "The text color", "Default": "empty"},
		{"Attribute": "Text", "Type": "string", "Description": "The example code", "Default": "empty"},
		{"Attribute": "Class", "Type": "string", "Description": "The other classes to apply to the card", "Default": "empty"},
	},
	Class: "bg-secondary",
}

var tableCode = templdais.TableAttrs{
	TableName: "templdais.LineCode",
	Columns:   []string{"Attribute", "Type", "Description", "Default"},
	Rows: []map[string]string{
		{"Attribute": "Prefix", "Type": "string", "Description": "The prefix of the code", "Default": "empty"},
		{"Attribute": "Brand", "Type": "string", "Description": "The text color", "Default": "empty"},
		{"Attribute": "Text", "Type": "string", "Description": "The example code", "Default": "empty"},
		{"Attribute": "Size", "Type": "string", "Description": "The size of the code", "Default": "empty"},
		{"Attribute": "Class", "Type": "string", "Description": "The other classes to apply to the card", "Default": "empty"},
	},
	Class: "bg-secondary",
}

var codes = []templdais.LineCode{
	{Prefix: "$", Brand: "success", Text: "go run .", Class: "w-[200px]"},
	{Prefix: "$", Brand: "error", Text: "npm i -D daisyui", Class: "w-[400px]"},
	{Prefix: "$", Brand: "warning", Text: "npx tailwindcss init -p", Class: "w-[400px]"},
}

func PreviewCode() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<article><section class=\"ml-10 lg:ml-0 bg-secondary rounded-lg p-10 flex justify-center items-start gap-x-4\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templdais.Code(templdais.CodeAttrs{Prefix: "$", Brand: "success", Text: "go run .", Class: "w-[200px]"}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templdais.Code(templdais.CodeAttrs{LineCode: codes, Class: "w-[400px]"}).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</section><section class=\" w-full flex justify-end !p-3\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Var2 := templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
			templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
			templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
			if !templ_7745c5c3_IsBuffer {
				defer func() {
					templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
					if templ_7745c5c3_Err == nil {
						templ_7745c5c3_Err = templ_7745c5c3_BufErr
					}
				}()
			}
			ctx = templ.InitializeContext(ctx)
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("View on DaisyUI <i class=\"fas fa-external-link-alt ml-2\"></i>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			return templ_7745c5c3_Err
		})
		templ_7745c5c3_Err = templdais.Button(templdais.ButtonAttrs{Type: "anchor", Link: templ.SafeURL("https://daisyui.com/components/accordion"), Outline: true, Brand: "accent", Class: "!px-2"}, templ.Attributes{"target": "_blank"}).Render(templ.WithChildren(ctx, templ_7745c5c3_Var2), templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</section></article>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func CodeComponentCode() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var3 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var3 == nil {
			templ_7745c5c3_Var3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<article><h2 class=\"text-lg font-bold\">Inline code</h2><h3 class=\"text-md italic font-semilight\">Single line</h3><pre data-prismjs-copy=\"Copy\" data-prismjs-copy-success=\"Success!\"><code class=\"language-go overflow-x-auto w-fit\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(
			`@templdais.Code(templdais.CodeAttrs{Prefix: "$", Brand: "success", Text: "go run .", Class: "w-[200px]"})
`)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/pages/code.templ`, Line: 71, Col: 1}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</code></pre><h3 class=\"text-md italic font-semilight\">Multiple lines</h3><pre data-prismjs-copy=\"Copy\" data-prismjs-copy-success=\"Success!\"><code class=\"language-go overflow-x-auto w-fit\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var5 string
		templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(
			`@templdais.Code(templdais.CodeAttrs{LineCode: templdais.LineCode{
	{Prefix: "$", Brand: "success", Text: "go run ."},
	{Prefix: "$", Brand: "error", Text: "npm i -D daisyui"},
	{Prefix: "$", Brand: "warning", Text: "npx tailwindcss init -p"},
}, Class: "w-[400px]"})
`)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/pages/code.templ`, Line: 86, Col: 1}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</code></pre><h2 class=\"text-lg font-bold\">Fragmented code</h2><h3 class=\"text-md italic font-semilight\">Single line</h3><pre class=\"w-full\" data-prismjs-copy=\"Copy\" data-prismjs-copy-success=\"Success!\"><code class=\"language-go w-full\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var6 string
		templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(
			`var attrs = templdais.CodeAttrs{
	Prefix: "$",
	Brand: "success",
	Text: "go run .",
	Class: "w-[200px]",
}

templ MyComponent() {
	@templdais.Code(attrs)
}
`)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/pages/code.templ`, Line: 108, Col: 1}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</code></pre><h3 class=\"text-md italic font-semilight\">Multiple lines</h3><pre class=\"w-full\" data-prismjs-copy=\"Copy\" data-prismjs-copy-success=\"Success!\"><code class=\"language-go w-full\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var7 string
		templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(
			`var codes = []templdais.LineCode{
	{Prefix: "$", Brand: "success", Text: "go run .", Class: "w-[200px]"},
	{Prefix: "$", Brand: "error", Text: "npm i -D daisyui", Class: "w-[400px]"},
	{Prefix: "$", Brand: "warning", Text: "npx tailwindcss init -p", Class: "w-[400px]"},
}

var attrs = templdais.CodeAttrs{
	Prefix: "$",
	Brand: "success",
	Text: "go run .",
	Class: "w-[200px]",
	LineCode: codes,
}

templ MyComponent() {
	@templdais.Code(attrs)
}
`)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/pages/code.templ`, Line: 136, Col: 1}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</code></pre></article>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func AttributesCode() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var8 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var8 == nil {
			templ_7745c5c3_Var8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<article><h2 class=\"text-md italic font-semilight\">templdais.LineCode</h2>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templdais.Table(tableCode).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</article><article class=\"my-3\"><h2 class=\"text-md italic font-semilight\">templdais.CodeAttrs</h2>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templdais.Table(tableCodeAttrs).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</article>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

func CodePage() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var9 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var9 == nil {
			templ_7745c5c3_Var9 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"hidden w-[400px]\"></div><section class=\"p-4 w-full\"><h1 class=\"text-2xl font-bold my-2\">Code Component</h1>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templdais.Accordion(accordionCode).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p class=\"hidden text-primary text-secondary text-info text-warning text-error text-success\"></p></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}
