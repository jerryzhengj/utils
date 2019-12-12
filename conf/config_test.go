package conf

import (
	"github.com/BurntSushi/toml"
	. "gopkg.in/check.v1"
	"os"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{
	LayOneSimple string `toml:"layOneSimple"`
	LayOneComplex LayTwo `toml:"layOneComplex"`
}


type LayTwo struct{
	LayTwoSimple string `toml:"layTwoSimple"`
	LayTwoSimple2 int `toml:"layTwoSimple2"`
	LayTwoComplex LayThird `toml:"layTwoComplex"`
}

type LayThird struct{
	LayThirdSimple int `toml:"layThirdSimple"`
	LayThirdSimple2 []string `toml:"layThirdSimple2"`
}

var _ = Suite(&MySuite{
	LayOneSimple:"layone",
	LayOneComplex: LayTwo{
		LayTwoSimple: "laytwo",
		LayTwoSimple2: 2,
		LayTwoComplex : LayThird{
			LayThirdSimple: 3,
			LayThirdSimple2: []string{"3-1","3-2","3-3"},
		},
	},

})

var fileName = "configtest.toml"

var confInst *configEnv


func (s *MySuite) TestLoad(c *C) {
	_, err :=confInst.Load()
	if err != nil{
		c.Fail()
	}
}

func (s *MySuite) TestRefresh(c *C) {
   err := confInst.Refresh()
	if err != nil{
		c.Fail()
	}
}

func (s *MySuite)TestGetCurrentConfig(c *C) {
	confInst.Load()
    conf,err := confInst.GetCurrentConfig()
	if err != nil{
		c.Fail()
	}

   obtained := conf.(*MySuite)
   c.Assert(obtained.LayOneSimple,Equals,"layone")
   c.Assert(obtained.LayOneComplex.LayTwoSimple2,Equals,2)
   c.Assert(obtained.LayOneComplex.LayTwoComplex.LayThirdSimple2[0],Equals,"3-1")
}

func (s *MySuite)TestSave(c *C) {
	confInst.Load()
    err :=confInst.Save()
    c.Assert(err,Equals,nil)
}

func (s *MySuite) SetUpSuite(c *C) {
	f, err := os.OpenFile(fileName, os.O_RDWR | os.O_CREATE|os.O_TRUNC, 0666)
	if err == nil {
		defer f.Close()
		err=toml.NewEncoder(f).Encode(s)
	}
	confInst =New(fileName,&MySuite{})
}

func (s *MySuite) TearDownSuite(c *C) {
	confInst = nil
	os.Remove(fileName)
}