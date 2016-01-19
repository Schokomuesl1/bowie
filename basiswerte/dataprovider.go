package basiswerte

import (
    "encoding/json"
    "io/ioutil"
)

var AlleEigenschaften map[string]string
var AlleSpezies map[string]SpeziesType
var AlleTalente []TalentType
var AlleKulturen map[string]KulturType

func init() {
    file, _ := ioutil.ReadFile("regeln/eigenschaften.json")
    AlleEigenschaften = make(map[string]string)
    json.Unmarshal([]byte(string(file)), &AlleEigenschaften)

    file2, _ := ioutil.ReadFile("regeln/spezies.json")
    speziesTmp := make([]SpeziesType, 0)
    json.Unmarshal([]byte(string(file2)), &speziesTmp)
    AlleSpezies = make(map[string]SpeziesType)
    for _, v := range speziesTmp {
        AlleSpezies[v.Name] = v
    }
    file3, _ := ioutil.ReadFile("regeln/talente.json")
    AlleTalente = make([]TalentType, 0)
    json.Unmarshal([]byte(string(file3)), &AlleTalente)
    file4, _ := ioutil.ReadFile("regeln/kulturen.json")
    kulturTmp := make([]KulturType, 0)
    AlleKulturen = make(map[string]KulturType)
    json.Unmarshal([]byte(string(file4)), &kulturTmp)
    for _, v := range kulturTmp {
        AlleKulturen[v.Name] = v
    }
}

type TalentType struct {
    Name      string
    Kategorie string
    Probe     [3]string
    Belastung string
    Kosten    string
}

type EigenschaftenModSpezies struct {
    Eigenschaft []string
    Mod         int
}

type SpeziesType struct {
    Name                       string
    LE                         int
    SK                         int
    ZK                         int
    GS                         int
    EigenschaftsModifikationen []EigenschaftenModSpezies
    Vorteile                   []string
    Nachteile                  []string
    AP                         int
}

type ModPair struct {
    Id    string
    Value int
}

type KulturType struct {
    Name    string
    Talente []ModPair
    Kosten  int
}
