package jobsearch

import (
	"log"
	"proto-course/protogen/basic"
	"proto-course/protogen/first"
	"proto-course/protogen/jobsearch"

	"google.golang.org/protobuf/encoding/protojson"
)

func JobSearchSample() {
	job := &jobsearch.JobSoftware{
		JobSoftwareId: 1,
		Application: &basic.Application{
			Version:   "1.0.0",
			Name:      "Proto software",
			Platforms: []string{"Windows", "Linux"},
		},
	}

	jsonBytes, err := protojson.Marshal(job)
	if err != nil {
		log.Fatal("error on marshal JobSoftware:", err)
	}

	log.Println("Software:", string(jsonBytes))
}

func JobSearchCandidate() {
	candidate := &jobsearch.JobCandidate{
		JobCandidateId: 1,
		Application: &first.Application{
			ApplicationId:     2,
			ApplicantFullName: "Víctor Reis",
			Phone:             "123456789",
			Email:             "victor@reis.com",
		},
	}

	jsonBytes, err := protojson.Marshal(candidate)
	if err != nil {
		log.Fatal("error on marshal JobCandidate:", err)
	}

	log.Println("Candidate:", string(jsonBytes))
}
