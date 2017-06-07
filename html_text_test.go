package main

import (
	"bytes"
	"fmt"
	"github.com/qri-io/go-diff/diffmatchpatch"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestHtmlTextContent(t *testing.T) {
	cases := []struct {
		testHtml   string
		expectText string
		err        error
	}{
		{"epa.gov.html", "epa.gov.txt", nil},
	}

	for i, c := range cases {
		in, err := ioutil.ReadFile(filepath.Join("test_files", c.testHtml))
		if err != nil {
			t.Errorf("case %d error reading test html file: %s", i, err.Error())
			continue
		}

		outbytes, err := ioutil.ReadFile(filepath.Join("test_files", c.expectText))
		if err != nil {
			t.Errorf("case %d error reading test output text file: %s", i, err.Error())
			continue
		}
		// out := string(outbytes)

		// create a fake response from the file
		got, err := HtmlTextContent(&http.Response{
			Request: httptest.NewRequest("GET", "http://testhtmltextcontent.com", nil),
			Body:    ioutil.NopCloser(bytes.NewBuffer(in)),
		})
		if err != c.err {
			t.Errorf("case %d error mismatch: %s != %s", i, c.err, err)
		}

		// fmt.Println(got)
		// ioutil.WriteFile("test_files/epa.guv.txt", outbytes, os.ModePerm)

		if got != output {
			// t.Errorf("case %d text output mismatch", i)

			filename := fmt.Sprintf("errors-%d.html", i)
			dmp := diffmatchpatch.New()
			ioutil.WriteFile(filename, []byte(dmp.DiffPrettyHtml(dmp.DiffMain(got, output, false))), os.ModePerm)
			t.Errorf("case %d text output mismatch diff of text written to: %s", i, filename)
		}
	}
}

const output = `Jump to main content US EPA
  Search

  Search





   Search Search Contact Us Contact Us Share




































































     Brownfields Cleanups172 communities will receive Brownfields grants totaling $56.8 M. The funds will aid under-served and economically disadvantaged communities across the nation.Learn more about the grants and communities.
$8 Million More for Water InfrastructureA bill signed by President Trump provides additional funds that will allow EPA to lend about $1.5 billion for water infrastructure projects nationwide.Read the news release


Brownfields Job TrainingThe Environmental Workforce Development and Job Training Program has trained more than 16,000 people for jobs in the environmental protection field. Today we’ve selected 14 organizations to continue this effort.Read the news releaseLearn moreRead the blog
Back-to-Basics: Clifton Hill, MOAdministrator Scott Pruitt speaks with power plant workers, Missouri Electric Cooperative members and agriculture leaders about EPA’s Back-to-Basics agenda.Read the news releaseView videos and photos from the event
 Brownfields Cleanups172 communities will receive Brownfields grants totaling $56.8 M. The funds will aid under-served and economically disadvantaged communities across the nation.Learn more about the grants and communities.
 172 communities will receive Brownfields grants totaling $56.8 M. The funds will aid under-served and economically disadvantaged communities across the nation.Learn more about the grants and communities. Learn more about the grants and communities. $8 Million More for Water InfrastructureA bill signed by President Trump provides additional funds that will allow EPA to lend about $1.5 billion for water infrastructure projects nationwide.Read the news release


 A bill signed by President Trump provides additional funds that will allow EPA to lend about $1.5 billion for water infrastructure projects nationwide.Read the news release Read the news release Brownfields Job TrainingThe Environmental Workforce Development and Job Training Program has trained more than 16,000 people for jobs in the environmental protection field. Today we’ve selected 14 organizations to continue this effort.Read the news releaseLearn moreRead the blog
 The Environmental Workforce Development and Job Training Program has trained more than 16,000 people for jobs in the environmental protection field. Today we’ve selected 14 organizations to continue this effort.Read the news releaseLearn moreRead the blog Read the news release Learn more Read the blog Back-to-Basics: Clifton Hill, MOAdministrator Scott Pruitt speaks with power plant workers, Missouri Electric Cooperative members and agriculture leaders about EPA’s Back-to-Basics agenda.Read the news releaseView videos and photos from the event
 Administrator Scott Pruitt speaks with power plant workers, Missouri Electric Cooperative members and agriculture leaders about EPA’s Back-to-Basics agenda.Read the news releaseView videos and photos from the event Back-to-Basics Read the news release View videos and photos from the event Administrator Pruitt traveled to Clifton Hill, Mo. to discuss EPA’s Back-to-Basics agenda. Watch the video. Back-to-Basics More EPA Videos More EPA Videos
    Grant for Water Quality Protection in South Carolina


    Protecting Water Quality in Belmont, Mass


    Wonder Farm Fined for Pesticide, Worker Protection Violations


    Portland, OR Student Wins EPA Award at Science Competition


    Grant for Water Quality Protection in South Carolina
   Grant for Water Quality Protection in South Carolina
    Protecting Water Quality in Belmont, Mass
   Protecting Water Quality in Belmont, Mass
    Wonder Farm Fined for Pesticide, Worker Protection Violations
   Wonder Farm Fined for Pesticide, Worker Protection Violations
    Portland, OR Student Wins EPA Award at Science Competition
   Portland, OR Student Wins EPA Award at Science Competition More news releases More news releases Tweet to @EPA >
  Learn about conditions where you live.
 Learn about conditions where you live. Go to EPA's page about your state.- choose -AlabamaAlaskaAmerican SamoaArizonaArkansasCaliforniaColoradoConnecticutDelawareDistrict of ColumbiaFloridaGeorgiaGuamHawaiiIdahoIllinoisIndianaIowaKansasKentuckyLouisianaMaineMarylandMassachusettsMichiganMinnesotaMississippiMissouriMontanaNebraskaNevadaNew HampshireNew JerseyNew MexicoNew YorkNorth CarolinaNorth DakotaN Mariana IslandsOhioOklahomaOregonPennsylvaniaPuerto RicoRhode IslandSouth CarolinaSouth DakotaTennesseeTexasTrust TerritoriesUtahVermontVirgin IslandsVirginiaWashingtonWest VirginiaWisconsinWyoming   Go to EPA's page about your state. - choose -AlabamaAlaskaAmerican SamoaArizonaArkansasCaliforniaColoradoConnecticutDelawareDistrict of ColumbiaFloridaGeorgiaGuamHawaiiIdahoIllinoisIndianaIowaKansasKentuckyLouisianaMaineMarylandMassachusettsMichiganMinnesotaMississippiMissouriMontanaNebraskaNevadaNew HampshireNew JerseyNew MexicoNew YorkNorth CarolinaNorth DakotaN Mariana IslandsOhioOklahomaOregonPennsylvaniaPuerto RicoRhode IslandSouth CarolinaSouth DakotaTennesseeTexasTrust TerritoriesUtahVermontVirgin IslandsVirginiaWashingtonWest VirginiaWisconsinWyoming - choose - Alabama Alaska American Samoa Arizona Arkansas California Colorado Connecticut Delaware District of Columbia Florida Georgia Guam Hawaii Idaho Illinois Indiana Iowa Kansas Kentucky Louisiana Maine Maryland Massachusetts Michigan Minnesota Mississippi Missouri Montana Nebraska Nevada New Hampshire New Jersey New Mexico New York North Carolina North Dakota N Mariana Islands Ohio Oklahoma Oregon Pennsylvania Puerto Rico Rhode Island South Carolina South Dakota Tennessee Texas Trust Territories Utah Vermont Virgin Islands Virginia Washington West Virginia Wisconsin Wyoming Get updates on topics you choose.   Get updates on topics you choose. Español | 中文: 繁體版  | 中文: 简体版 Tiếng Việt  | 한국어 Español 中文: 繁體版  中文: 简体版  Tiếng Việt  한국어 "I seek to listen, learn, and lead with you to address these issues we face as a nation." "I seek to listen, learn, and lead with you to address these issues we face as a nation." Administrator's bio Follow the Administrator on Twitter Contact Us to ask a question, provide feedback, or report a problem. Contact Us Environmental TopicsAir
Bed Bugs
Chemicals and Toxics
Environmental Information by Location
Greener Living
Health
Land, Waste, and Cleanup
Lead
Mold
Pesticides
Radon
Science
Water
A-Z Index

Laws & RegulationsBy Business Sector
By Topic
Compliance
Enforcement
Laws and Executive Orders
Policy and Guidance
Regulations

About EPAEPA Administrator
Current Leadership
Organization Chart
Staff Directory
Planning, Budget and Results
Jobs and Internships
Headquarters Offices
Regional Offices
Labs and Research Centers

 Environmental TopicsAir
Bed Bugs
Chemicals and Toxics
Environmental Information by Location
Greener Living
Health
Land, Waste, and Cleanup
Lead
Mold
Pesticides
Radon
Science
Water
A-Z Index
 Environmental Topics Air
Bed Bugs
Chemicals and Toxics
Environmental Information by Location
Greener Living
Health
Land, Waste, and Cleanup
Lead
Mold
Pesticides
Radon
Science
Water
A-Z Index
 Air Air Bed Bugs Bed Bugs Chemicals and Toxics Chemicals and Toxics Environmental Information by Location Environmental Information by Location Greener Living Greener Living Health Health Land, Waste, and Cleanup Land, Waste, and Cleanup Lead Lead Mold Mold Pesticides Pesticides Radon Radon Science Science Water Water A-Z Index A-Z Index Laws & RegulationsBy Business Sector
By Topic
Compliance
Enforcement
Laws and Executive Orders
Policy and Guidance
Regulations
 Laws & Regulations By Business Sector
By Topic
Compliance
Enforcement
Laws and Executive Orders
Policy and Guidance
Regulations
 By Business Sector By Business Sector By Topic By Topic Compliance Compliance Enforcement Enforcement Laws and Executive Orders Laws and Executive Orders Policy and Guidance Policy and Guidance Regulations Regulations About EPAEPA Administrator
Current Leadership
Organization Chart
Staff Directory
Planning, Budget and Results
Jobs and Internships
Headquarters Offices
Regional Offices
Labs and Research Centers
 About EPA EPA Administrator
Current Leadership
Organization Chart
Staff Directory
Planning, Budget and Results
Jobs and Internships
Headquarters Offices
Regional Offices
Labs and Research Centers
 EPA Administrator EPA Administrator Current Leadership Current Leadership Organization Chart Organization Chart Staff Directory Staff Directory Planning, Budget and Results Planning, Budget and Results Jobs and Internships Jobs and Internships Headquarters Offices Headquarters Offices Regional Offices Regional Offices Labs and Research Centers Labs and Research Centers
      Accessibility
      EPA Administrator
      Budget & Performance
      Contracting
      Grants
      January 19, 2017 Web Snapshot
      No FEAR Act Data
      Privacy
      Privacy and Security Notice
     Accessibility Accessibility EPA Administrator EPA Administrator Budget & Performance Budget & Performance Contracting Contracting Grants Grants January 19, 2017 Web Snapshot January 19, 2017 Web Snapshot No FEAR Act Data No FEAR Act Data Privacy Privacy Privacy and Security Notice Privacy and Security Notice
      Data.gov
      Inspector General
      Jobs
      Newsroom
      Open Government
      Regulations.gov
      Subscribe
      USA.gov
      White House
     Data.gov Data.gov Inspector General Inspector General Jobs Jobs Newsroom Newsroom Open Government Open Government Regulations.gov Regulations.gov Subscribe Subscribe USA.gov USA.gov White House White House
      Contact Us
      Hotlines
      FOIA Requests
      Frequent Questions
     Contact Us Contact Us Hotlines Hotlines FOIA Requests FOIA Requests Frequent Questions Frequent Questions
      Facebook
      Twitter
      YouTube
      Flickr
      Instagram
     Facebook Facebook Twitter Twitter YouTube YouTube Flickr Flickr Instagram Instagram Last updated on June 2, 2017`