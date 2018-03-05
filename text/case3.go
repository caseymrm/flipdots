package text

func getCase3JSON() []byte {
	return []byte(`
{
	"Name": "CaseThree",
	"Height": 3,
	"Width": 3,
	"Characters": [
		{
			"Character": "0",
			"Bitmap": [
				"---",
				"---",
				"---"
			]
		},
		{
			"Character": "1",
			"Bitmap": [
				"00-",
				"-0-",
				"000"
			]
		},
		{
			"Character": "2",
			"Bitmap": [
				"00-",
				"-0-",
				"-00"
			]
		},
		{
			"Character": "3",
			"Bitmap": [
				"000",
				"-00",
				"000"
			]
		},
		{
			"Character": "4",
			"Bitmap": [
				"0-0",
				"000",
				"--0"
			]
		},
		{
			"Character": "5",
			"Bitmap": [
				"-00",
				"-0-",
				"00-"
			]
		},
		{
			"Character": "6",
			"Bitmap": [
				"0--",
				"000",
				"000"
			]
		},
		{
			"Character": "7",
			"Bitmap": [
				"000",
				"-0-",
				"0--"
			]
		},
		{
			"Character": "8",
			"Bitmap": [
				"000",
				"000",
				"000"
			]
		},
		{
			"Character": "9",
			"Bitmap": [
				"000",
				"000",
				"--0"
			]
		},
		{
			"Character": "?",
			"Bitmap": [
				"000",
				"-00",
				"-0-"
			]
		},
		{
			"Character": " ",
			"Bitmap": [
				"---",
				"---",
				"---"
			]
		}
	]
}
`)
}
