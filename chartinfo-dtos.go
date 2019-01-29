package main

import "encoding/xml"

type CARMC struct {
	XMLName xml.Name `xml:"ARMC,omitempty" json:"ARMC,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CAscendant struct {
	XMLName xml.Name `xml:"Ascendant,omitempty" json:"Ascendant,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CCo_dash_Ascendant1 struct {
	XMLName xml.Name `xml:"Co-Ascendant1,omitempty" json:"Co-Ascendant1,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CCo_dash_Ascendant2 struct {
	XMLName xml.Name `xml:"Co-Ascendant2,omitempty" json:"Co-Ascendant2,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CConjunction struct {
	XMLName xml.Name `xml:"Conjunction,omitempty" json:"Conjunction,omitempty"`
	Attrbody1 string`xml:"body1,attr"  json:",omitempty"`
	Attrbody2 string`xml:"body2,attr"  json:",omitempty"`
	Attrdegree1 string`xml:"degree1,attr"  json:",omitempty"`
	Attrdegree2 string`xml:"degree2,attr"  json:",omitempty"`
}

type CEarth struct {
	XMLName xml.Name `xml:"Earth,omitempty" json:"Earth,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrdist string`xml:"dist,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrretrograde string`xml:"retrograde,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CEquatorialAscendant struct {
	XMLName xml.Name `xml:"EquatorialAscendant,omitempty" json:"EquatorialAscendant,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CHouse struct {
	XMLName xml.Name `xml:"House,omitempty" json:"House,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrnumber string`xml:"number,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CJupiter struct {
	XMLName xml.Name `xml:"Jupiter,omitempty" json:"Jupiter,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrdist string`xml:"dist,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrretrograde string`xml:"retrograde,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CMC struct {
	XMLName xml.Name `xml:"MC,omitempty" json:"MC,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CMars struct {
	XMLName xml.Name `xml:"Mars,omitempty" json:"Mars,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrdist string`xml:"dist,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrretrograde string`xml:"retrograde,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CMeanNode struct {
	XMLName xml.Name `xml:"MeanNode,omitempty" json:"MeanNode,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrdist string`xml:"dist,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrretrograde string`xml:"retrograde,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CMeanSouthNode struct {
	XMLName xml.Name `xml:"MeanSouthNode,omitempty" json:"MeanSouthNode,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrdist string`xml:"dist,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrretrograde string`xml:"retrograde,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CMercury struct {
	XMLName xml.Name `xml:"Mercury,omitempty" json:"Mercury,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrdist string`xml:"dist,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrretrograde string`xml:"retrograde,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CMoon struct {
	XMLName xml.Name `xml:"Moon,omitempty" json:"Moon,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrdist string`xml:"dist,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrretrograde string`xml:"retrograde,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CNeptune struct {
	XMLName xml.Name `xml:"Neptune,omitempty" json:"Neptune,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrdist string`xml:"dist,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrretrograde string`xml:"retrograde,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type COpposition struct {
	XMLName xml.Name `xml:"Opposition,omitempty" json:"Opposition,omitempty"`
	Attrbody1 string`xml:"body1,attr"  json:",omitempty"`
	Attrbody2 string`xml:"body2,attr"  json:",omitempty"`
	Attrdegree1 string`xml:"degree1,attr"  json:",omitempty"`
	Attrdegree2 string`xml:"degree2,attr"  json:",omitempty"`
}

type CPluto struct {
	XMLName xml.Name `xml:"Pluto,omitempty" json:"Pluto,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrdist string`xml:"dist,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrretrograde string`xml:"retrograde,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CPolarAscendant struct {
	XMLName xml.Name `xml:"PolarAscendant,omitempty" json:"PolarAscendant,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CSaturn struct {
	XMLName xml.Name `xml:"Saturn,omitempty" json:"Saturn,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrdist string`xml:"dist,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrretrograde string`xml:"retrograde,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CSextile struct {
	XMLName xml.Name `xml:"Sextile,omitempty" json:"Sextile,omitempty"`
	Attrbody1 string`xml:"body1,attr"  json:",omitempty"`
	Attrbody2 string`xml:"body2,attr"  json:",omitempty"`
	Attrdegree1 string`xml:"degree1,attr"  json:",omitempty"`
	Attrdegree2 string`xml:"degree2,attr"  json:",omitempty"`
}

type CSquare struct {
	XMLName xml.Name `xml:"Square,omitempty" json:"Square,omitempty"`
	Attrbody1 string`xml:"body1,attr"  json:",omitempty"`
	Attrbody2 string`xml:"body2,attr"  json:",omitempty"`
	Attrdegree1 string`xml:"degree1,attr"  json:",omitempty"`
	Attrdegree2 string`xml:"degree2,attr"  json:",omitempty"`
}

type CSun struct {
	XMLName xml.Name `xml:"Sun,omitempty" json:"Sun,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrdist string`xml:"dist,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrretrograde string`xml:"retrograde,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CTrine struct {
	XMLName xml.Name `xml:"Trine,omitempty" json:"Trine,omitempty"`
	Attrbody1 string`xml:"body1,attr"  json:",omitempty"`
	Attrbody2 string`xml:"body2,attr"  json:",omitempty"`
	Attrdegree1 string`xml:"degree1,attr"  json:",omitempty"`
	Attrdegree2 string`xml:"degree2,attr"  json:",omitempty"`
}

type CUranus struct {
	XMLName xml.Name `xml:"Uranus,omitempty" json:"Uranus,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrdist string`xml:"dist,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrretrograde string`xml:"retrograde,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CVenus struct {
	XMLName xml.Name `xml:"Venus,omitempty" json:"Venus,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrdist string`xml:"dist,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrretrograde string`xml:"retrograde,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type CVertex struct {
	XMLName xml.Name `xml:"Vertex,omitempty" json:"Vertex,omitempty"`
	Attrdegree string`xml:"degree,attr"  json:",omitempty"`
	Attrdegree_ut string`xml:"degree_ut,attr"  json:",omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrsign string`xml:"sign,attr"  json:",omitempty"`
	Attrsign_name string`xml:"sign_name,attr"  json:",omitempty"`
}

type Cascmcs struct {
	XMLName xml.Name `xml:"ascmcs,omitempty" json:"ascmcs,omitempty"`
	CARMC *CARMC `xml:"ARMC,omitempty" json:"ARMC,omitempty"`
	CAscendant *CAscendant `xml:"Ascendant,omitempty" json:"Ascendant,omitempty"`
	CCo_dash_Ascendant1 *CCo_dash_Ascendant1 `xml:"Co-Ascendant1,omitempty" json:"Co-Ascendant1,omitempty"`
	CCo_dash_Ascendant2 *CCo_dash_Ascendant2 `xml:"Co-Ascendant2,omitempty" json:"Co-Ascendant2,omitempty"`
	CEquatorialAscendant *CEquatorialAscendant `xml:"EquatorialAscendant,omitempty" json:"EquatorialAscendant,omitempty"`
	CMC *CMC `xml:"MC,omitempty" json:"MC,omitempty"`
	CPolarAscendant *CPolarAscendant `xml:"PolarAscendant,omitempty" json:"PolarAscendant,omitempty"`
	CVertex *CVertex `xml:"Vertex,omitempty" json:"Vertex,omitempty"`
}

type Caspects struct {
	XMLName xml.Name `xml:"aspects,omitempty" json:"aspects,omitempty"`
	CConjunction []*CConjunction `xml:"Conjunction,omitempty" json:"Conjunction,omitempty"`
	COpposition []*COpposition `xml:"Opposition,omitempty" json:"Opposition,omitempty"`
	CSextile []*CSextile `xml:"Sextile,omitempty" json:"Sextile,omitempty"`
	CSquare []*CSquare `xml:"Square,omitempty" json:"Square,omitempty"`
	CTrine []*CTrine `xml:"Trine,omitempty" json:"Trine,omitempty"`
}

type Cbodies struct {
	XMLName xml.Name `xml:"bodies,omitempty" json:"bodies,omitempty"`
	CEarth *CEarth `xml:"Earth,omitempty" json:"Earth,omitempty"`
	CJupiter *CJupiter `xml:"Jupiter,omitempty" json:"Jupiter,omitempty"`
	CMars *CMars `xml:"Mars,omitempty" json:"Mars,omitempty"`
	CMeanNode *CMeanNode `xml:"MeanNode,omitempty" json:"MeanNode,omitempty"`
	CMeanSouthNode *CMeanSouthNode `xml:"MeanSouthNode,omitempty" json:"MeanSouthNode,omitempty"`
	CMercury *CMercury `xml:"Mercury,omitempty" json:"Mercury,omitempty"`
	CMoon *CMoon `xml:"Moon,omitempty" json:"Moon,omitempty"`
	CNeptune *CNeptune `xml:"Neptune,omitempty" json:"Neptune,omitempty"`
	CPluto *CPluto `xml:"Pluto,omitempty" json:"Pluto,omitempty"`
	CSaturn *CSaturn `xml:"Saturn,omitempty" json:"Saturn,omitempty"`
	CSun *CSun `xml:"Sun,omitempty" json:"Sun,omitempty"`
	CUranus *CUranus `xml:"Uranus,omitempty" json:"Uranus,omitempty"`
	CVenus *CVenus `xml:"Venus,omitempty" json:"Venus,omitempty"`
}

type Cchartinfo struct {
	XMLName xml.Name `xml:"chartinfo,omitempty" json:"chartinfo,omitempty"`
	Attrcity string`xml:"city,attr"  json:",omitempty"`
	Attrday string`xml:"day,attr"  json:",omitempty"`
	Attrdisplay string`xml:"display,attr"  json:",omitempty"`
	Attrhsys string`xml:"hsys,attr"  json:",omitempty"`
	Attrlat string`xml:"lat,attr"  json:",omitempty"`
	Attrlon string`xml:"lon,attr"  json:",omitempty"`
	Attrmonth string`xml:"month,attr"  json:",omitempty"`
	Attrname string`xml:"name,attr"  json:",omitempty"`
	Attrtime string`xml:"time,attr"  json:",omitempty"`
	Attryear string`xml:"year,attr"  json:",omitempty"`
	Cascmcs *Cascmcs `xml:"ascmcs,omitempty" json:"ascmcs,omitempty"`
	Caspects *Caspects `xml:"aspects,omitempty" json:"aspects,omitempty"`
	Cbodies *Cbodies `xml:"bodies,omitempty" json:"bodies,omitempty"`
	Chouses *Chouses `xml:"houses,omitempty" json:"houses,omitempty"`
}

type Chouses struct {
	XMLName xml.Name `xml:"houses,omitempty" json:"houses,omitempty"`
	CHouse []*CHouse `xml:"House,omitempty" json:"House,omitempty"`
}