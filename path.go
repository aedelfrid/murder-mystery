package main

import (
	//"fmt"
	//"github.com/manifoldco/promptui"
)

type path struct {
	id string
	scene []string
	prompt struct {
		question string
		answers []string
		pathids []string
	}
}

// func importPaths

func servePath(p path) {

}

/*
	function to serve path
		function to serve scene
			function to prompt

			|
			\/
	feed path to new serve function?
*/