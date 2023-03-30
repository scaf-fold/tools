package xconvert

import (
	"fmt"
	"net/url"
	"os"
	"sync"

	_ "github.com/lib/pq"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type Model struct {
	Models []ModelItemConfig `yaml:"Model"`
}

type ModelItemConfig struct {
	Conn  string            `yaml:"Dsn"`
	Table map[string]string `yaml:"Table"`
}

type ModelConverter struct {
	targetFile string
}

func NewModelConverter(configFile string) *ModelConverter {
	return &ModelConverter{
		targetFile: configFile,
	}
}

func (m *ModelConverter) Converter() (*Model, error) {
	content, err := os.ReadFile(m.targetFile)
	if err != nil {
		return nil, err
	}
	model := &Model{}
	err = yaml.Unmarshal(content, model)
	if err != nil {
		return nil, err
	}
	return model, err
}

func (m *ModelConverter) Gen() *Model {
	ms, err := m.Converter()
	if err != nil {
		panic(err)
	}
	wg := &sync.WaitGroup{}
	for _, item := range ms.Models {
		wg.Add(1)
		go func(conf ModelItemConfig) {
			db, err := gorm.Open(postgres.New(postgres.Config{
				DSN: conf.Conn,
			}))
			defer func() {
				d, err := db.DB()
				if err != nil {
					fmt.Println(err)
				} else {
					err := d.Close()
					if err != nil {
						return
					}
				}
				wg.Done()
			}()
			if err != nil {
				panic(err)
			}
			baseUrl, err := url.Parse(conf.Conn)
			if err != nil {
				panic(err)
			}
			outPath := fmt.Sprintf("./db/%s/gen/query", baseUrl.Path)
			g := gen.NewGenerator(gen.Config{
				OutPath: outPath,
			})
			g.UseDB(db)
			g.ApplyBasic(func(g *gen.Generator, table map[string]string) []interface{} {
				m := make([]interface{}, 0)
				for k, v := range table {
					m = append(m, g.GenerateModelAs(k, v))
				}
				return m
			}(g, conf.Table)...)
			g.Execute()

		}(item)
	}
	wg.Wait()
	return ms
}
