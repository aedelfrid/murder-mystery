# Murder Mystery

Murder Mystery is a CYA style game using promptui and maybe possibly a game engine.

in this game you assume the role of Billiam Wurdoch; a detective in the employ of the Woronto Sheriff
where you solve crimes.

## Planning

programmatic prompt

prompt import

point system?

game engine implementation

### prompt import

json format?

    [
        {
            "path": {
                "id": "path0",
                "scene": [
                    "lorem ipsum",
                    "lorem ipsum",
                    "lorem ipsum",
                    "lorem ipsum",
                    "lorem ipsum",
                    "lorem ipsum",
                    "lorem ipsum",
                ],
                "prompt": {
                    "question":"lorem ipsum?",
                    "answers": ["yes", "no"],
                    "pathids": ["path1", "path2"]
                }
            }

        }
    ]


paths are imported -> recc'd to use a flowchart to create paths
this way branching paths can be created, adding and removing 
programmatically over time.

a readable and reusable pattern

### programmatic prompt

prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
			"Saturday", "Sunday"},
	}

_, result, err := prompt.Run()

if err != nil {
	fmt.Printf("Prompt failed %v\n", err)
    return
}

fmt.Printf("You choose %q\n", result)



#### context

for

### game engine

