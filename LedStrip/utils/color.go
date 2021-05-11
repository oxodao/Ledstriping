package utils

func GetHexFromBoard(boardColor string) string {
	return "#" + boardColor[2:4] + boardColor[6:8] + boardColor[4:6]
}

func GetBoardFromHex(hexColor string) string {
	return "0x" + hexColor[1:3] + hexColor[5:7] + hexColor[3:5]
}