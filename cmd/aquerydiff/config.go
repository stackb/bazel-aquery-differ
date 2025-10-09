package main

type config struct {
	target           string
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
