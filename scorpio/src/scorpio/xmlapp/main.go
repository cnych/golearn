package main

import (
	"encoding/xml"
	"fmt"
)

var (
	xmlContent = `<?xml version="1.0" encoding="UTF-8" standalone="no" ?>
	<xml_result>
    <read_sentence lan="en" type="study" version="">
        <rec_paper>
            <read_chapter beg_pos="0" content="How are you?" end_pos="164" except_info="0" is_rejected="false" total_score="4.480
392" word_count="3">
                <sentence beg_pos="0" content="how are you" end_pos="164" index="0" total_score="4.480392" word_count="3">
                    <word beg_pos="48" content="how" dp_message="0" end_pos="73" global_index="0" index="0" total_score="3.722479
">
                        <syll beg_pos="48" content="hh aw" end_pos="73" syll_score="3.722479">
                            <phone beg_pos="48" content="hh" dp_message="0" end_pos="68" />
                            <phone beg_pos="68" content="aw" dp_message="0" end_pos="73" />
                        </syll>
                    </word>
                    <word beg_pos="73" content="are" dp_message="0" end_pos="92" global_index="1" index="1" total_score="4.969362
">
                        <syll beg_pos="73" content="aa" end_pos="92" syll_score="4.969362">
                            <phone beg_pos="73" content="aa" dp_message="0" end_pos="92" />
                        </syll>
                    </word>
                    <word beg_pos="92" content="you" dp_message="0" end_pos="128" global_index="2" index="2" total_score="4.91563
7">
                        <syll beg_pos="92" content="y uw" end_pos="128" syll_score="4.915637">
                            <phone beg_pos="92" content="y" dp_message="0" end_pos="100" />
                            <phone beg_pos="100" content="uw" dp_message="0" end_pos="128" />
                        </syll>
                    </word>
                </sentence>
            </read_chapter>
        </rec_paper>
    </read_sentence>
</xml_result>`
	xmlContent1 = `<?xml version="1.0" encoding="UTF-8"?>
	<person>
  <name>Luann Van Houten</name>
  <addresses>
      <address type="secondary">
          <street>321 MadeUp Lane</street>
          <city>Shelbyville</city>
      </address>
      <address type="primary">
          <street>123 Fake St</street>
          <city>Springfield</city>
      </address>
  </addresses>
</person>`
)

type Person struct {
	Name      string `xml:"name"`
	Addresses []struct {
		Street string `xml:"street"`
		City   string `xml:"city"`
		Type   string `xml:"type,attr"`
	} `xml:"addresses>address"`
}

type XMLResult struct {
	ReadSentence struct {
		Language string `xml:"lan,attr"`
		Category string `xml:"type,attr"`
		RecPaper struct {
			ReadChapter struct {
				BegPos     string `xml:"beg_pos,attr"`
				EndPos     string `xml:"end_pos,attr"`
				Content    string `xml:"content,attr"`
				ExceptInfo string `xml:"except_info,attr"`
				IsRejected string `xml:"is_rejected,attr"`
				TotalScore string `xml:"total_score,attr"`
				WordCount  string `xml:"word_count,attr"`
				Sentences  []struct {
					BegPos     string `xml:"beg_pos,attr"`
					EndPos     string `xml:"end_pos,attr"`
					Content    string `xml:"content,attr"`
					Index      string `xml:"index,attr"`
					TotalScore string `xml:"total_score,attr"`
					WordCount  string `xml:"word_count,attr"`
					Words      []struct {
						BegPos      string `xml:"beg_pos,attr"`
						EndPos      string `xml:"end_pos,attr"`
						Content     string `xml:"content,attr"`
						DpMessage   string `xml:"dp_message,attr"`
						GlobalIndex string `xml:"global_index,attr"`
						Index       string `xml:"index,attr"`
						TotalScore  string `xml:"total_score,attr"`
					} `xml:"word"`
				} `xml:"sentence"`
			} `xml:"read_chapter"`
		} `xml:"rec_paper"`
	} `xml:"read_sentence"`
}

func main() {
	var per Person
	err := xml.Unmarshal([]byte(xmlContent1), &per)
	if err != nil {
		fmt.Printf("err=%v\n", err)
		return
	}
	fmt.Printf("per=%v\n", per)

	var iseRepo XMLResult
	err = xml.Unmarshal([]byte(xmlContent), &iseRepo)
	if err != nil {
		fmt.Printf("err=%v\n", err)
		return
	}
	//fmt.Printf("iseRepo=%v\n", iseRepo)
	fmt.Printf("sentences=%v\n", iseRepo.ReadSentence.RecPaper.ReadChapter.Sentences)
}
