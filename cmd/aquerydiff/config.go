package main

type config struct {
	beforeFile       string
	afterFile        string
	reportDir        string
	matchingStrategy string
	port             string
	unidiff          bool
	cmpdiff          bool
	serve            bool
	open             bool
}
