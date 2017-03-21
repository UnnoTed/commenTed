package commenTed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tmpl := []byte(`
  // c:remove
  package hello
  // c:end
  
  // c:remove
  package ayy
  // c:end
  `)

	etmpl := `
  
  package hello

  
  
  package ayy

  `

	tmpl = Parse(tmpl, false)
	assert.Equal(t, etmpl, string(tmpl))

	/////////////////////////

	tmpl = []byte(`
  // c:remove
  // package hello
  // c:end

  // c:remove
  // package ayy
  // c:end
  `)

	etmpl = `
  
  package hello


  
  package ayy

  `

	tmpl = Parse(tmpl, false)
	assert.Equal(t, etmpl, string(tmpl))

	//////////////////////////

	tmpl = []byte(`
  // c:remove
  package hello // c:too
  // c:end
  
  // c:remove
  package ayy
  // c:end
  `)

	etmpl = `
  


  
  
  package ayy

  `

	tmpl = Parse(tmpl, false)
	assert.Equal(t, etmpl, string(tmpl))

}

func TestComplex(t *testing.T) {
	tmpl := []byte(`
// c:remove
package auth // c:too

// package {{.Package}}
// c:end

// c:remove
// {{range .Fields}}
// 	 {{.Name}} {{.Type}} {{.Tags}}
// {{end}}
// c:end
  `)

	etmpl := `



package {{.Package}}



{{range .Fields}}
{{.Name}} {{.Type}} {{.Tags}}
{{end}}

  `

	tmpl = Parse(tmpl, false)
	assert.Equal(t, etmpl, string(tmpl))
}

func TestReplacer(t *testing.T) {
	tmpl := []byte(`
func (u *User) FindByNAME(data string) error {
  // c:replace:up [User|Users] - [data string|id hide.Int64] - [NAME|ID]
}
  `)

	etmpl := `
func (u *Users) FindByID(id hide.Int64) error {

}
  `

	tmpl = ParseReplace(tmpl, "[", "]", true)
	assert.Equal(t, etmpl, string(tmpl))
}
