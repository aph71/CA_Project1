/*******  TO DO    *******/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/***********************************************************************/
/*****************************B TYPE INSTRUCTIONS***********************/
/***********************************************************************/

type BranchInstruction struct {
	Opcode     string
	OffsetStr  string
	OffsetInt  int
	LineNumber int
}

func NewBranchInstruction(opcode string, offsetStr string, offsetInt int, lineNumber int) *BranchInstruction {
	return &BranchInstruction{
		Opcode:     opcode,
		OffsetStr:  offsetStr,
		OffsetInt:  offsetInt,
		LineNumber: lineNumber,
	}
}

func parseBranchInstruction(binaryInstruction string, lineNumber int) *BranchInstruction {
	opcode := binaryInstruction[:6]
	offsetStr := binaryInstruction[6:32]
	offsetInt, _ := twosComplement(offsetStr)
	return NewBranchInstruction(opcode, offsetStr, offsetInt, lineNumber)
}

/***********************************************************************/
/*****************************CB TYPE INSTRUCTIONS***********************/
/***********************************************************************/

type cBranchInstruction struct {
	Opcode      string
	OffsetStr   string
	RegistryStr string
	RegistryInt int
	OffsetInt   int
	LineNumber  int
}

func NewcBranchInstruction(opcode string, offsetStr string, registryStr string,
	registryInt int, offsetInt int, lineNumber int) *cBranchInstruction {
	return &cBranchInstruction{
		Opcode:      opcode,
		OffsetStr:   offsetStr,
		RegistryStr: registryStr,
		RegistryInt: registryInt,
		OffsetInt:   offsetInt,
		LineNumber:  lineNumber,
	}
}

func parsecBranchInstruction(binaryInstruction string, lineNumber int) *cBranchInstruction {
	opcode := binaryInstruction[:8]
	offsetStr := binaryInstruction[8:27]
	registryStr := binaryInstruction[27:32]
	registryInt, _ := strconv.ParseInt(registryStr, 2, 32)
	offsetInt, _ := twosComplement(offsetStr)
	return NewcBranchInstruction(opcode, offsetStr, registryStr, int(registryInt), offsetInt, lineNumber)
}

/***********************************************************************/
/*****************************IM TYPE INSTRUCTIONS***********************/
/***********************************************************************/
type imTypeInstruction struct {
	Opcode      string
	FieldValue  string
	DestStr     string
	Op2         string
	Destination int
	LineNumber  int
	FieldInt    uint
	BitPattern  int
}

func NewImTypeInstruction(opcode string, op2 string, destStr string, fieldValue string,
	bitPattern int, dest int, fieldInt uint, lineNumber int) *imTypeInstruction {
	return &imTypeInstruction{
		Opcode:      opcode,
		Op2:         op2,
		Destination: dest,
		FieldInt:    fieldInt,
		LineNumber:  lineNumber,
		FieldValue:  fieldValue,
		DestStr:     destStr,
		BitPattern:  bitPattern,
	}
}

func parseImTypeInstruction(binaryInstruction string, lineNumber int) *imTypeInstruction {
	opcode := binaryInstruction[:9]
	op2 := binaryInstruction[9:11]
	fieldValue := binaryInstruction[11:27]
	destStr := binaryInstruction[27:32]
	dest, _ := strconv.ParseInt(destStr, 2, 32)
	fieldInt, _ := strconv.ParseInt(fieldValue, 2, 32)
	bitPattern := 0
	if op2 == "00" {
		bitPattern = 0
	} else if op2 == "01" {
		bitPattern = 16
	} else if op2 == "10" {
		bitPattern = 32
	} else if op2 == "11" {
		bitPattern = 48
	}
	return NewImTypeInstruction(opcode, op2, destStr, fieldValue, bitPattern, int(dest), uint(fieldInt), lineNumber)
}

/***********************************************************************/
/*****************************I TYPE INSTRUCTIONS***********************/
/***********************************************************************/
type iTypeInstruction struct {
	Opcode      string
	Immediate   string
	Src1Str     string
	DestStr     string
	Destination int
	Source1     int
	LineNumber  int
	ImmedInt    int
}

func NewITypeInstruction(opcode string, immediate string,
	src1Str string, destStr string, dest int, src1 int, immedInt int, lineNumber int) *iTypeInstruction {
	return &iTypeInstruction{
		Opcode:      opcode,
		Destination: dest,
		Source1:     src1,
		ImmedInt:    immedInt,
		LineNumber:  lineNumber,
		Immediate:   immediate,
		Src1Str:     src1Str,
		DestStr:     destStr,
	}
}

func parseITypeInstruction(binaryInstruction string, lineNumber int) *iTypeInstruction {
	opcode := binaryInstruction[:10]
	immediate := binaryInstruction[10:22]
	src1Str := binaryInstruction[22:27]
	destStr := binaryInstruction[27:32]
	dest, _ := strconv.ParseInt(destStr, 2, 32)
	src1, _ := strconv.ParseInt(src1Str, 2, 32)
	immedInt, _ := strconv.ParseInt(immediate, 2, 32)
	return NewITypeInstruction(opcode, immediate, src1Str, destStr, int(dest), int(src1), int(immedInt), lineNumber)
}

/***********************************************************************/
/*****************************R TYPE INSTRUCTIONS***********************/
/***********************************************************************/
type rTypeInstruction struct {
	Opcode      string
	Destination int
	Source1     int
	Source2     int
	LineNumber  int
	ShamtInt    int
	Shamt       string
	Src2Str     string
	Src1Str     string
	DestStr     string
}

func NewrTypeInstruction(opcode string, shamt string, src2Str string,
	src1Str string, destStr string, dest int, src1 int, src2 int, shamtInt int, lineNumber int) *rTypeInstruction {
	return &rTypeInstruction{
		Opcode:      opcode,
		Destination: dest,
		Source1:     src1,
		Source2:     src2,
		ShamtInt:    shamtInt,
		LineNumber:  lineNumber,
		Shamt:       shamt,
		Src2Str:     src2Str,
		Src1Str:     src1Str,
		DestStr:     destStr,
	}
}

func parserTypeInstruction(binaryInstruction string, lineNumber int) *rTypeInstruction {
	opcode := binaryInstruction[:11]
	src2Str := binaryInstruction[11:16]
	shamt := binaryInstruction[16:22]
	src1Str := binaryInstruction[22:27]
	destStr := binaryInstruction[27:32]
	dest, _ := strconv.ParseInt(destStr, 2, 32)
	src1, _ := strconv.ParseInt(src1Str, 2, 32)
	src2, _ := strconv.ParseInt(src2Str, 2, 32)
	shamtInt, _ := strconv.ParseInt(shamt, 2, 32)
	return NewrTypeInstruction(opcode, shamt, src2Str, src1Str, destStr, int(dest), int(src1), int(src2), int(shamtInt), lineNumber)
}

/*******************************************************************/
/**************************D TYPE INSTRUCTIONS***********************/
/*******************************************************************/
type dTypeInstruction struct {
	Opcode     string
	AddressStr string
	Op2        string
	Src2Str    string
	Src1Str    string
	AddressInt int
	Source1    int
	Source2    int
	LineNumber int
}

func NewDTypeInstruction(opcode string, addressStr string, op2 string,
	src1Str string, src2Str string, addressInt int, src1 int, src2 int, lineNumber int) *dTypeInstruction {
	return &dTypeInstruction{
		Opcode:     opcode,
		AddressStr: addressStr,
		Source1:    src1,
		Source2:    src2,
		LineNumber: lineNumber,
		Op2:        op2,
		Src2Str:    src2Str,
		Src1Str:    src1Str,
		AddressInt: addressInt,
	}
}

func parseDTypeInstruction(binaryInstruction string, lineNumber int) *dTypeInstruction {
	opcode := binaryInstruction[0:11]
	addressStr := binaryInstruction[11:20]
	op2 := binaryInstruction[20:22]
	src2Str := binaryInstruction[22:27]
	src1Str := binaryInstruction[27:32]
	addressInt, _ := strconv.ParseInt(addressStr, 2, 32)
	src1, _ := strconv.ParseInt(src1Str, 2, 32)
	src2, _ := strconv.ParseInt(src2Str, 2, 32)
	return NewDTypeInstruction(opcode, addressStr, op2, src1Str, src2Str, int(addressInt), int(src1), int(src2), lineNumber)
}

/*********************************************************/
/*********************************************************/
/***********      MAIN DRIVER FUNCTION   ****************/
/*********************************************************/
/*********************************************************/

func readAndProcessInstructions(binaryInstruction string, lineNumber int, outputFile *string, output *string) {

	/************B TYPE INSTRUCTIONS****************/
	switch binaryInstruction[:6] {
	case "000101":
		branchInst := parseBranchInstruction(binaryInstruction, lineNumber)
		*output += fmt.Sprintf("%.6s %.26s\t\t%.1d\tB\t#%d\n",
			branchInst.Opcode, branchInst.OffsetStr, lineNumber, branchInst.OffsetInt)
		//fmt.Printf(output)
		//writeToFile(*outputFile, output)
	default:
		/************CB TYPE INSTRUCTIONS****************/

		switch binaryInstruction[:8] {
		case "10110100":
			cbTypeinst := parsecBranchInstruction(binaryInstruction, lineNumber)
			*output += fmt.Sprintf("%.8s %.19s %.5s\t\t%.1d\tCBZ\tR%.1d, #%.1d\n",
				cbTypeinst.Opcode, cbTypeinst.OffsetStr, cbTypeinst.RegistryStr, lineNumber, cbTypeinst.RegistryInt,
				cbTypeinst.OffsetInt)
			//fmt.Printf(output)
			//writeToFile(*outputFile, output)

		case "10110101":
			cbTypeinst := parsecBranchInstruction(binaryInstruction, lineNumber)
			*output += fmt.Sprintf("%.8s %.19s %.5s\t%.1d\tCBNZ\tR%.1d, #%.1d\n",
				cbTypeinst.Opcode, cbTypeinst.OffsetStr, cbTypeinst.RegistryStr, lineNumber, cbTypeinst.RegistryInt,
				cbTypeinst.OffsetInt)
			//fmt.Printf(output)
			//writeToFile(*outputFile, output)
		default:
			/************IM TYPE INSTRUCTIONS****************/
			switch binaryInstruction[:9] {
			case "110100101":
				imTypeinst := parseImTypeInstruction(binaryInstruction, lineNumber)
				*output += fmt.Sprintf("%.9s %.2s %.16s %.5s\t%.1d\tMOVZ\tR%.1d, %.1d, LSL %.1d\n",
					imTypeinst.Opcode, imTypeinst.Op2, imTypeinst.FieldValue, imTypeinst.DestStr, lineNumber,
					imTypeinst.Destination, imTypeinst.FieldInt, imTypeinst.BitPattern)
				//fmt.Printf(output)
				//writeToFile(*outputFile, output)
			case "111100101":
				imTypeinst := parseImTypeInstruction(binaryInstruction, lineNumber)
				*output += fmt.Sprintf("%.9s %.2s %.16s %.5s\t%.1d\tMOVK\tR%.1d, %.1d, LSL %.1d\n",
					imTypeinst.Opcode, imTypeinst.Op2, imTypeinst.FieldValue, imTypeinst.DestStr, lineNumber,
					imTypeinst.Destination, imTypeinst.FieldInt, imTypeinst.BitPattern)
				//fmt.Printf(output)
				//writeToFile(*outputFile, output)
			default:
				/***********I TYPE INSTRUCTIONS******************/

				/***********SUBI TYPE INSTRUCTIONS******************/
				switch binaryInstruction[:10] {
				case "1101000100":
					iTypeInst := parseITypeInstruction(binaryInstruction, lineNumber)
					*output += fmt.Sprintf("%.10s %.12s %.5s %.5s\t\t%.1d\tSUBI\tR%.1d, R%.1d, #%.1d\n",
						iTypeInst.Opcode, iTypeInst.Immediate, iTypeInst.Src1Str, iTypeInst.DestStr, lineNumber,
						iTypeInst.Destination, iTypeInst.Source1, iTypeInst.ImmedInt)
					//fmt.Printf(output)
					//writeToFile(*outputFile, output)
					/***********ADDI TYPE INSTRUCTIONS******************/
				case "1001000100":
					iTypeInst := parseITypeInstruction(binaryInstruction, lineNumber)
					*output += fmt.Sprintf("%.10s %.12s %.5s %.5s \t%.1d ADDI\tR%.1d, R%.1d, #%.1d\n",
						iTypeInst.Opcode, iTypeInst.Immediate, iTypeInst.Src1Str, iTypeInst.DestStr, lineNumber,
						iTypeInst.Source1, iTypeInst.Destination, iTypeInst.ImmedInt)
					//fmt.Printf(output)
					//writeToFile(*outputFile, output)
				default:
					/*************D TYPE INSTRUCTIONS****************/
					switch binaryInstruction[:11] {
					//*******STUR******//
					case "11111000000":
						dTypeInst := parseDTypeInstruction(binaryInstruction, lineNumber)
						*output += fmt.Sprintf("%.11s %.9s %.2s %.5s %.5s\t%.1d\tSTUR\tR%.1d, [R%.1d, #%.1d]\n",
							dTypeInst.Opcode, dTypeInst.AddressStr, dTypeInst.Op2, dTypeInst.Src2Str,
							dTypeInst.Src1Str, lineNumber, dTypeInst.Source1, dTypeInst.Source2, dTypeInst.AddressInt)
						//fmt.Printf(output)
						//writeToFile(*outputFile, output)
						//*******LDUR******//
					case "11111000010": //
						dTypeInst := parseDTypeInstruction(binaryInstruction, lineNumber)
						*output += fmt.Sprintf("%.11s %.9s %.2s %.5s %.5s\t%.1d\tLDUR\tR%.1d, [R%.1d, #%.1d]\n",
							dTypeInst.Opcode, dTypeInst.AddressStr, dTypeInst.Op2, dTypeInst.Src2Str,
							dTypeInst.Src1Str, lineNumber, dTypeInst.Source1, dTypeInst.Source2, dTypeInst.AddressInt)
						//fmt.Printf(output)
						//writeToFile(*outputFile, output)
						/*******************RTYPE INSTRUCTIONS******************/
					default:
						switch binaryInstruction[:11] {
						//*******ADD******//
						case "10001011000":
							rTypeInst := parserTypeInstruction(binaryInstruction, lineNumber)
							*output += fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s\t%.1d\tADD\tR%.1d, R%.1d, R%.1d\n",
								rTypeInst.Opcode, rTypeInst.Src2Str, rTypeInst.Shamt, rTypeInst.Src1Str, rTypeInst.DestStr, lineNumber,
								rTypeInst.Destination, rTypeInst.Source1, rTypeInst.Source2)
							//fmt.Printf(output)
							//writeToFile(*outputFile, output)
							//*******AND******//
						case "10001010000":
							rTypeInst := parserTypeInstruction(binaryInstruction, lineNumber)
							*output += fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s\t%.1d\tAND\tR%.1d, R%.1d, R%.1d\n",
								rTypeInst.Opcode, rTypeInst.Src2Str, rTypeInst.Shamt, rTypeInst.Src1Str, rTypeInst.DestStr, lineNumber,
								rTypeInst.Destination, rTypeInst.Source1, rTypeInst.Source2)
							//fmt.Printf(output)
							//writeToFile(*outputFile, output)
							//*******ORR******//
						case "10101010000":
							rTypeInst := parserTypeInstruction(binaryInstruction, lineNumber)
							*output += fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s\t%.1d\tORR\tR%.1d, R%.1d, R%.1d\n",
								rTypeInst.Opcode, rTypeInst.Src2Str, rTypeInst.Shamt, rTypeInst.Src1Str, rTypeInst.DestStr, lineNumber,
								rTypeInst.Destination, rTypeInst.Source1, rTypeInst.Source2)
							//fmt.Printf(output)
							//writeToFile(*outputFile, output)
							//****************SUB***************//
						case "11001011000":
							rTypeInst := parserTypeInstruction(binaryInstruction, lineNumber)
							*output += fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s\t%.1d\tSUB\tR%.1d, R%.1d, R%.1d\n",
								rTypeInst.Opcode, rTypeInst.Src2Str, rTypeInst.Shamt, rTypeInst.Src1Str, rTypeInst.DestStr, lineNumber,
								rTypeInst.Destination, rTypeInst.Source1, rTypeInst.Source2)
							//fmt.Printf(output)
							//writeToFile(*outputFile, output)
							//***************EOR****************//
						case "11101010000":
							rTypeInst := parserTypeInstruction(binaryInstruction, lineNumber)
							*output += fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s\t%.1d\tEOR\tR%.1d, R%.1d, R%.1d\n",
								rTypeInst.Opcode, rTypeInst.Src2Str, rTypeInst.Shamt, rTypeInst.Src1Str, rTypeInst.DestStr, lineNumber,
								rTypeInst.Destination, rTypeInst.Source1, rTypeInst.Source2)
							//fmt.Printf(output)
							//writeToFile(*outputFile, output)
							//***********ASR*****************//
						case "11010011100":
							rTypeInst := parserTypeInstruction(binaryInstruction, lineNumber)
							*output += fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s\t%.1d\tASR\tR%.1d, R%.1d, #%.1d\n",
								rTypeInst.Opcode, rTypeInst.Src2Str, rTypeInst.Shamt, rTypeInst.Src1Str, rTypeInst.DestStr, lineNumber,
								rTypeInst.Destination, rTypeInst.Source1, rTypeInst.ShamtInt)
							//fmt.Printf(output)
							//writeToFile(*outputFile, output)
							//**************LSL******************//
						case "11010011011":
							rTypeInst := parserTypeInstruction(binaryInstruction, lineNumber)
							*output += fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s\t%.1d\tLSL\tR%.1d, R%.1d, #%.1d\n",
								rTypeInst.Opcode, rTypeInst.Src2Str, rTypeInst.Shamt, rTypeInst.Src1Str, rTypeInst.DestStr, lineNumber,
								rTypeInst.Destination, rTypeInst.Source1, rTypeInst.ShamtInt)
							//fmt.Printf(output)
							//writeToFile(*outputFile, output)
							//**************LSR*******************//
						case "11010011010":
							rTypeInst := parserTypeInstruction(binaryInstruction, lineNumber)
							*output += fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s\t%.1d\tLSR\tR%.1d, R%.1d, #%.1d\n",
								rTypeInst.Opcode, rTypeInst.Src2Str, rTypeInst.Shamt, rTypeInst.Src1Str, rTypeInst.DestStr, lineNumber,
								rTypeInst.Destination, rTypeInst.Source1, rTypeInst.ShamtInt)
							//fmt.Printf(output)
							//writeToFile(*outputFile, output)
						default:
							switch binaryInstruction[:32] {
							case "11111110110111101111111111100111":
								*output += fmt.Sprintf("%.8s %.3s %.5s %.5s %.5s %.6s\t%.1d\tBREAK\n",
									binaryInstruction[:8], binaryInstruction[8:11], binaryInstruction[11:16], binaryInstruction[16:21],
									binaryInstruction[21:26], binaryInstruction[26:32], lineNumber)
								//fmt.Printf(output)
								//writeToFile(*outputFile, output)
							case "00000000000000000000000000000000":
								*output += fmt.Sprintf("%.8s %.3s %.5s %.5s %.5s %.6s\t%.1d\tNOP\n",
									binaryInstruction[:8], binaryInstruction[8:11], binaryInstruction[11:16], binaryInstruction[16:21],
									binaryInstruction[21:26], binaryInstruction[26:32], lineNumber)
								//fmt.Printf(output)
								//writeToFile(*outputFile, output)
							default:
								calc, err := twosComplement(binaryInstruction)
								calcString := strconv.Itoa(calc)
								*output += fmt.Sprintf("%.8s %.3s %.5s %.5s %.5s %.5s\t%.1d\t%.5s\n",
									binaryInstruction[:8], binaryInstruction[8:11], binaryInstruction[11:16], binaryInstruction[16:21],
									binaryInstruction[21:26], binaryInstruction[26:32], lineNumber, calcString)
								if err != nil {
								}
								//fmt.Printf(output)
								//writeToFile(*outputFile, output)
							}
						}
					}
				}

			}
		}
	}
}

/*********************************************************/
/*********************************************************/
/***********      MATH       FUNCTIONS    ****************/
/*********************************************************/
/*********************************************************/

/***********************2s COMPLEMENT**********************/

func twosComplement(binaryStr string) (int, error) {
	var negFlag bool //Check for leading one (negative number)
	if binaryStr[0:1] != "1" {
		negFlag = false
		result2, err := binaryToInteger(binaryStr)
		return result2, err
	} else {
		negFlag = true
	}
	inverted := ""

	// Invert each bit individually
	for _, bit := range binaryStr {
		if bit == '0' {
			inverted += "1"
		} else {
			inverted += "0"
		}
	}
	// Trim leading zeros
	binaryStr = strings.TrimLeft(inverted, "0")
	result := ""
	// Carry for addition
	carry := 1
	// Work through string from right to left
	for i := len(binaryStr) - 1; i >= 0; i-- {
		bit := int(binaryStr[i] - '0')
		sum := bit + carry
		// Update result and carry
		result = strconv.Itoa(sum%2) + result
		carry = sum / 2
	}
	// Add leftover carry to left
	if carry == 1 {
		result = "1" + result
	}
	// Add zeroes to return to original length
	for len(result) < len(binaryStr) {
		result = "0" + result
	}
	result2, err := binaryToInteger(result)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if negFlag == true {
		result2 = -result2
	}
	return result2, nil
}

/***************BINARY TO INTEGER CONVERTER*************/
func binaryToInteger(binary string) (int, error) {
	result, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		return 0, err
	}
	return int(result), nil
}

/************   MAIN   ****************/

func main() {
	// String flag with default value "input.txt"
	inputFile := flag.String("i", "input.txt", "Enter input file")

	// String flag with default value "team16_out_dis.txt"
	outputFile := flag.String("o", "team16_out", "Enter desired name for the output file")

	// Enable command-line parsing
	flag.Parse()
	//simulator output file
	outputSimFile := *outputFile + "_sim.txt"
	//disassembler output file
	*outputFile += "_dis.txt"
	openInputFile, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer func(inputFile *os.File) {
		err := inputFile.Close()
		if err != nil {
		}
	}(openInputFile)

	// Create a reader to read from the input file
	reader := bufio.NewReader(openInputFile)
	//EDITED FROM WRITE TO FILE FUNCTION

	writeOutputFile, err := os.OpenFile(*outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(writeOutputFile)
	lineNumber := 96
	output := ""
	for {

		// Create a new buffer for each loop
		buffer := make([]byte, 32)
		// Read 32 characters into new buffer
		var bytesRead int
		for bytesRead < 32 {
			char, err := reader.ReadByte()
			if err != nil {
				if err == io.EOF {
					// End of file reached, break the loop
					break
				} else {
					fmt.Println("Error reading from file:", err)
					return
				}
			}
			// Skip newline characters and empty spaces
			if char == '\n' {
				continue
			}
			if unicode.IsSpace(rune(char)) {
				continue
			}
			buffer[bytesRead] = char
			bytesRead++
		}
		if bytesRead < 32 {
			// End of file reached, break the loop
			break
		}
		// Convert  to string
		binaryInstruction := string(buffer[:bytesRead])
		readAndProcessInstructions(binaryInstruction, lineNumber, outputFile, &output)
		lineNumber += 4
	}
	//writing to file
	_, err = writeOutputFile.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		// file.Close()
		return
	}
	err = writeOutputFile.Close()
	if err != nil {
		return
	}

	//Initializing each variable used
	instructionAddress := 96
	cycle := 1
	outputSim := ""

	R0 := 0
	R1 := 0
	R2 := 0
	R3 := 0
	R4 := 0
	R5 := 0
	R6 := 0
	R7 := 0
	R8 := 0
	R9 := 0
	R10 := 0
	R11 := 0
	R12 := 0
	R13 := 0
	R14 := 0
	R15 := 0
	R16 := 0
	R17 := 0
	R18 := 0
	R19 := 0
	R20 := 0
	R21 := 0
	R22 := 0
	R23 := 0
	R24 := 0
	R25 := 0
	R26 := 0
	R27 := 0
	R28 := 0
	R29 := 0
	R30 := 0
	R31 := 0

	//instructionString := ""
	openOutputFile, err := os.Open(*outputFile)
	if err != nil {
		fmt.Println("Error opening output file:", err)
		return
	}
	defer func(OutputFile *os.File) {
		err := OutputFile.Close()
		if err != nil {
		}
	}(openOutputFile)

	//WRITE TO SIMULATOR FILE
	simScanner := bufio.NewScanner(openOutputFile)
	simScanner.Split(bufio.ScanLines)
	writeOutputSimFile, err := os.OpenFile(outputSimFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(writeOutputSimFile)
	//write simulator format
	for instructionAddress < lineNumber {
		simScanner.Scan()
		newLine := simScanner.Text()
		length := len(newLine)
		opcode := newLine[:strings.IndexByte(newLine, ' ')]
		instructionString := ""

		if len(opcode) == 6 {
			if instructionAddress < 100 {
				instructionString = newLine[38:length]
			} else {
				instructionString = newLine[39:length]
			}
			//
		}
		if len(opcode) == 8 {
			if instructionAddress < 100 {
				instructionString = newLine[39:length]
			} else {
				instructionString = newLine[40:length]
			}
			//THE BREAK CONDITION
			if opcode == "11111110" {
				instructionString = newLine[42:length]
			}
		}
		if len(opcode) == 9 {
			//TEST THIS ONE
			if instructionAddress < 100 {
				instructionString = newLine[41:length]
			} else {
				instructionString = newLine[42:length]
			}
		}
		//ADDI AND SUBI
		if len(opcode) == 10 {

			if instructionAddress < 100 {
				instructionString = newLine[40:length]
				
			} else {
				instructionString = newLine[41:length]
				
			}
			//TEST THIS ONE

			lineLength := len(instructionString)

			registerThree := instructionString[5:7]
			registerTwo := instructionString[9:11]

			immediateInt := instructionString[14:lineLength]
			immediate, _ := strconv.Atoi(immediateInt)

			count1 := 0
			count2 := 0
			for i := 0; i < 32; i++ {
				var registerNum = strconv.Itoa(i)
				registerNum = "R" + registerNum

				if registerNum == registerThree {
					count1 = i
				}
				if registerNum == registerTwo {
					count2 = i
				}
			}

			if instructionString[0:4] == "SUBI" {
				//SUBI
				if count1 == 0 {
					if count2 == 0 {
						R0 -= immediate
					}
					if count2 == 1 {
						R1 = R0 - immediate
					}
					if count2 == 2 {
						R2 = R0 - immediate
					}
					if count2 == 3 {
						R3 = R0 - immediate
					}
					if count2 == 4 {
						R4 = R0 - immediate
					}
					if count2 == 5 {
						R5 = R0 - immediate
					}
					if count2 == 6 {
						R6 = R0 - immediate
					}
					if count2 == 7 {
						R7 = R0 - immediate
					}
					if count2 == 8 {
						R8 = R0 - immediate
					}
					if count2 == 9 {
						R9 = R0 - immediate
					}
					if count2 == 10 {
						R10 = R0 - immediate
					}
					if count2 == 11 {
						R11 = R0 - immediate
					}
					if count2 == 12 {
						R12 = R0 - immediate
					}
					if count2 == 13 {
						R13 = R0 - immediate
					}
					if count2 == 14 {
						R14 = R0 - immediate
					}
					if count2 == 15 {
						R15 = R0 - immediate
					}
					if count2 == 16 {
						R16 = R0 - immediate
					}
					if count2 == 17 {
						R17 = R0 - immediate
					}
					if count2 == 18 {
						R18 = R0 - immediate
					}
					if count2 == 19 {
						R19 = R0 - immediate
					}
					if count2 == 20 {
						R20 = R0 - immediate
					}
					if count2 == 21 {
						R21 = R0 - immediate
					}
					if count2 == 22 {
						R22 = R0 - immediate
					}
					if count2 == 23 {
						R23 = R0 - immediate
					}
					if count2 == 24 {
						R24 = R0 - immediate
					}
					if count2 == 25 {
						R25 = R0 - immediate
					}
					if count2 == 26 {
						R26 = R0 - immediate
					}
					if count2 == 27 {
						R27 = R0 - immediate
					}
					if count2 == 28 {
						R28 = R0 - immediate
					}
					if count2 == 29 {
						R29 = R0 - immediate
					}
					if count2 == 30 {
						R30 = R0 - immediate
					}
					if count2 == 31 {
						R31 = R0 - immediate
					}
				}
				if count1 == 1 {
					if count2 == 0 {
						R0 = R1 - immediate
					}
					if count2 == 1 {
						R1 = R1 - immediate
					}
					if count2 == 2 {
						R2 = R1 - immediate
					}
					if count2 == 3 {
						R3 = R1 - immediate
					}
					if count2 == 4 {
						R4 = R1 - immediate
					}
					if count2 == 5 {
						R5 = R1 - immediate
					}
					if count2 == 6 {
						R6 = R1 - immediate
					}
					if count2 == 7 {
						R7 = R1 - immediate
					}
					if count2 == 8 {
						R8 = R1 - immediate
					}
					if count2 == 9 {
						R9 = R1 - immediate
					}
					if count2 == 10 {
						R10 = R1 - immediate
					}
					if count2 == 11 {
						R11 = R1 - immediate
					}
					if count2 == 12 {
						R12 = R1 - immediate
					}
					if count2 == 13 {
						R13 = R1 - immediate
					}
					if count2 == 14 {
						R14 = R1 - immediate
					}
					if count2 == 15 {
						R15 = R1 - immediate
					}
					if count2 == 16 {
						R16 = R1 - immediate
					}
					if count2 == 17 {
						R17 = R1 - immediate
					}
					if count2 == 18 {
						R18 = R1 - immediate
					}
					if count2 == 19 {
						R19 = R1 - immediate
					}
					if count2 == 20 {
						R20 = R1 - immediate
					}
					if count2 == 21 {
						R21 = R1 - immediate
					}
					if count2 == 22 {
						R22 = R1 - immediate
					}
					if count2 == 23 {
						R23 = R1 - immediate
					}
					if count2 == 24 {
						R24 = R1 - immediate
					}
					if count2 == 25 {
						R25 = R1 - immediate
					}
					if count2 == 26 {
						R26 = R1 - immediate
					}
					if count2 == 27 {
						R27 = R1 - immediate
					}
					if count2 == 28 {
						R28 = R1 - immediate
					}
					if count2 == 29 {
						R29 = R1 - immediate
					}
					if count2 == 30 {
						R30 = R1 - immediate
					}
					if count2 == 31 {
						R31 = R1 - immediate
					}
				}
				if count1 == 2 {
					if count2 == 0 {
						R0 = R2 - immediate
					}
					if count2 == 1 {
						R1 = R2 - immediate
					}
					if count2 == 2 {
						R2 = R2 - immediate
					}
					if count2 == 3 {
						R3 = R2 - immediate
					}
					if count2 == 4 {
						R4 = R2 - immediate
					}
					if count2 == 5 {
						R5 = R2 - immediate
					}
					if count2 == 6 {
						R6 = R2 - immediate
					}
					if count2 == 7 {
						R7 = R2 - immediate
					}
					if count2 == 8 {
						R8 = R2 - immediate
					}
					if count2 == 9 {
						R9 = R2 - immediate
					}
					if count2 == 10 {
						R10 = R2 - immediate
					}
					if count2 == 11 {
						R11 = R2 - immediate
					}
					if count2 == 12 {
						R12 = R2 - immediate
					}
					if count2 == 13 {
						R13 = R2 - immediate
					}
					if count2 == 14 {
						R14 = R2 - immediate
					}
					if count2 == 15 {
						R15 = R2 - immediate
					}
					if count2 == 16 {
						R16 = R2 - immediate
					}
					if count2 == 17 {
						R17 = R2 - immediate
					}
					if count2 == 18 {
						R18 = R2 - immediate
					}
					if count2 == 19 {
						R19 = R2 - immediate
					}
					if count2 == 20 {
						R20 = R2 - immediate
					}
					if count2 == 21 {
						R21 = R2 - immediate
					}
					if count2 == 22 {
						R22 = R2 - immediate
					}
					if count2 == 23 {
						R23 = R2 - immediate
					}
					if count2 == 24 {
						R24 = R2 - immediate
					}
					if count2 == 25 {
						R25 = R2 - immediate
					}
					if count2 == 26 {
						R26 = R2 - immediate
					}
					if count2 == 27 {
						R27 = R2 - immediate
					}
					if count2 == 28 {
						R28 = R2 - immediate
					}
					if count2 == 29 {
						R29 = R2 - immediate
					}
					if count2 == 30 {
						R30 = R2 - immediate
					}
					if count2 == 31 {
						R31 = R2 - immediate
					}
				}
				if count1 == 3 {
					if count2 == 0 {
						R0 = R3 - immediate
					}
					if count2 == 1 {
						R1 = R3 - immediate
					}
					if count2 == 2 {
						R2 = R3 - immediate
					}
					if count2 == 3 {
						R3 = R3 - immediate
					}
					if count2 == 4 {
						R4 = R3 - immediate
					}
					if count2 == 5 {
						R5 = R3 - immediate
					}
					if count2 == 6 {
						R6 = R3 - immediate
					}
					if count2 == 7 {
						R7 = R3 - immediate
					}
					if count2 == 8 {
						R8 = R3 - immediate
					}
					if count2 == 9 {
						R9 = R3 - immediate
					}
					if count2 == 10 {
						R10 = R3 - immediate
					}
					if count2 == 11 {
						R11 = R3 - immediate
					}
					if count2 == 12 {
						R12 = R3 - immediate
					}
					if count2 == 13 {
						R13 = R3 - immediate
					}
					if count2 == 14 {
						R14 = R3 - immediate
					}
					if count2 == 15 {
						R15 = R3 - immediate
					}
					if count2 == 16 {
						R16 = R3 - immediate
					}
					if count2 == 17 {
						R17 = R3 - immediate
					}
					if count2 == 18 {
						R18 = R3 - immediate
					}
					if count2 == 19 {
						R19 = R3 - immediate
					}
					if count2 == 20 {
						R20 = R3 - immediate
					}
					if count2 == 21 {
						R21 = R3 - immediate
					}
					if count2 == 22 {
						R22 = R3 - immediate
					}
					if count2 == 23 {
						R23 = R3 - immediate
					}
					if count2 == 24 {
						R24 = R3 - immediate
					}
					if count2 == 25 {
						R25 = R3 - immediate
					}
					if count2 == 26 {
						R26 = R3 - immediate
					}
					if count2 == 27 {
						R27 = R3 - immediate
					}
					if count2 == 28 {
						R28 = R3 - immediate
					}
					if count2 == 29 {
						R29 = R3 - immediate
					}
					if count2 == 30 {
						R30 = R3 - immediate
					}
					if count2 == 31 {
						R31 = R3 - immediate
					}
				}
				if count1 == 4 {
					if count2 == 0 {
						R0 = R4 - immediate
					}
					if count2 == 1 {
						R1 = R4 - immediate
					}
					if count2 == 2 {
						R2 = R4 - immediate
					}
					if count2 == 3 {
						R3 = R4 - immediate
					}
					if count2 == 4 {
						R4 = R4 - immediate
					}
					if count2 == 5 {
						R5 = R4 - immediate
					}
					if count2 == 6 {
						R6 = R4 - immediate
					}
					if count2 == 7 {
						R7 = R4 - immediate
					}
					if count2 == 8 {
						R8 = R4 - immediate
					}
					if count2 == 9 {
						R9 = R4 - immediate
					}
					if count2 == 10 {
						R10 = R4 - immediate
					}
					if count2 == 11 {
						R11 = R4 - immediate
					}
					if count2 == 12 {
						R12 = R4 - immediate
					}
					if count2 == 13 {
						R13 = R4 - immediate
					}
					if count2 == 14 {
						R14 = R4 - immediate
					}
					if count2 == 15 {
						R15 = R4 - immediate
					}
					if count2 == 16 {
						R16 = R4 - immediate
					}
					if count2 == 17 {
						R17 = R4 - immediate
					}
					if count2 == 18 {
						R18 = R4 - immediate
					}
					if count2 == 19 {
						R19 = R4 - immediate
					}
					if count2 == 20 {
						R20 = R4 - immediate
					}
					if count2 == 21 {
						R21 = R4 - immediate
					}
					if count2 == 22 {
						R22 = R4 - immediate
					}
					if count2 == 23 {
						R23 = R4 - immediate
					}
					if count2 == 24 {
						R24 = R4 - immediate
					}
					if count2 == 25 {
						R25 = R4 - immediate
					}
					if count2 == 26 {
						R26 = R4 - immediate
					}
					if count2 == 27 {
						R27 = R4 - immediate
					}
					if count2 == 28 {
						R28 = R4 - immediate
					}
					if count2 == 29 {
						R29 = R4 - immediate
					}
					if count2 == 30 {
						R30 = R4 - immediate
					}
					if count2 == 31 {
						R31 = R4 - immediate
					}
				}
				if count1 == 5 {
					if count2 == 0 {
						R0 = R5 - immediate
					}
					if count2 == 1 {
						R1 = R5 - immediate
					}
					if count2 == 2 {
						R2 = R5 - immediate
					}
					if count2 == 3 {
						R3 = R5 - immediate
					}
					if count2 == 4 {
						R4 = R5 - immediate
					}
					if count2 == 5 {
						R5 = R5 - immediate
					}
					if count2 == 6 {
						R6 = R5 - immediate
					}
					if count2 == 7 {
						R7 = R5 - immediate
					}
					if count2 == 8 {
						R8 = R5 - immediate
					}
					if count2 == 9 {
						R9 = R5 - immediate
					}
					if count2 == 10 {
						R10 = R5 - immediate
					}
					if count2 == 11 {
						R11 = R5 - immediate
					}
					if count2 == 12 {
						R12 = R5 - immediate
					}
					if count2 == 13 {
						R13 = R5 - immediate
					}
					if count2 == 14 {
						R14 = R5 - immediate
					}
					if count2 == 15 {
						R15 = R5 - immediate
					}
					if count2 == 16 {
						R16 = R5 - immediate
					}
					if count2 == 17 {
						R17 = R5 - immediate
					}
					if count2 == 18 {
						R18 = R5 - immediate
					}
					if count2 == 19 {
						R19 = R5 - immediate
					}
					if count2 == 20 {
						R20 = R5 - immediate
					}
					if count2 == 21 {
						R21 = R5 - immediate
					}
					if count2 == 22 {
						R22 = R5 - immediate
					}
					if count2 == 23 {
						R23 = R5 - immediate
					}
					if count2 == 24 {
						R24 = R5 - immediate
					}
					if count2 == 25 {
						R25 = R5 - immediate
					}
					if count2 == 26 {
						R26 = R5 - immediate
					}
					if count2 == 27 {
						R27 = R5 - immediate
					}
					if count2 == 28 {
						R28 = R5 - immediate
					}
					if count2 == 29 {
						R29 = R5 - immediate
					}
					if count2 == 30 {
						R30 = R5 - immediate
					}
					if count2 == 31 {
						R31 = R5 - immediate
					}
				}
				if count1 == 6 {
					if count2 == 0 {
						R0 = R6 - immediate
					}
					if count2 == 1 {
						R1 = R6 - immediate
					}
					if count2 == 2 {
						R2 = R6 - immediate
					}
					if count2 == 3 {
						R3 = R6 - immediate
					}
					if count2 == 4 {
						R4 = R6 - immediate
					}
					if count2 == 5 {
						R5 = R6 - immediate
					}
					if count2 == 6 {
						R6 = R6 - immediate
					}
					if count2 == 7 {
						R7 = R6 - immediate
					}
					if count2 == 8 {
						R8 = R6 - immediate
					}
					if count2 == 9 {
						R9 = R6 - immediate
					}
					if count2 == 10 {
						R10 = R6 - immediate
					}
					if count2 == 11 {
						R11 = R6 - immediate
					}
					if count2 == 12 {
						R12 = R6 - immediate
					}
					if count2 == 13 {
						R13 = R6 - immediate
					}
					if count2 == 14 {
						R14 = R6 - immediate
					}
					if count2 == 15 {
						R15 = R6 - immediate
					}
					if count2 == 16 {
						R16 = R6 - immediate
					}
					if count2 == 17 {
						R17 = R6 - immediate
					}
					if count2 == 18 {
						R18 = R6 - immediate
					}
					if count2 == 19 {
						R19 = R6 - immediate
					}
					if count2 == 20 {
						R20 = R6 - immediate
					}
					if count2 == 21 {
						R21 = R6 - immediate
					}
					if count2 == 22 {
						R22 = R6 - immediate
					}
					if count2 == 23 {
						R23 = R6 - immediate
					}
					if count2 == 24 {
						R24 = R6 - immediate
					}
					if count2 == 25 {
						R25 = R6 - immediate
					}
					if count2 == 26 {
						R26 = R6 - immediate
					}
					if count2 == 27 {
						R27 = R6 - immediate
					}
					if count2 == 28 {
						R28 = R6 - immediate
					}
					if count2 == 29 {
						R29 = R6 - immediate
					}
					if count2 == 30 {
						R30 = R6 - immediate
					}
					if count2 == 31 {
						R31 = R6 - immediate
					}
				}
				if count1 == 7 {
					if count2 == 0 {
						R0 = R7 - immediate
					}
					if count2 == 1 {
						R1 = R7 - immediate
					}
					if count2 == 2 {
						R2 = R7 - immediate
					}
					if count2 == 3 {
						R3 = R7 - immediate
					}
					if count2 == 4 {
						R4 = R7 - immediate
					}
					if count2 == 5 {
						R5 = R7 - immediate
					}
					if count2 == 6 {
						R6 = R7 - immediate
					}
					if count2 == 7 {
						R7 = R7 - immediate
					}
					if count2 == 8 {
						R8 = R7 - immediate
					}
					if count2 == 9 {
						R9 = R7 - immediate
					}
					if count2 == 10 {
						R10 = R7 - immediate
					}
					if count2 == 11 {
						R11 = R7 - immediate
					}
					if count2 == 12 {
						R12 = R7 - immediate
					}
					if count2 == 13 {
						R13 = R7 - immediate
					}
					if count2 == 14 {
						R14 = R7 - immediate
					}
					if count2 == 15 {
						R15 = R7 - immediate
					}
					if count2 == 16 {
						R16 = R7 - immediate
					}
					if count2 == 17 {
						R17 = R7 - immediate
					}
					if count2 == 18 {
						R18 = R7 - immediate
					}
					if count2 == 19 {
						R19 = R7 - immediate
					}
					if count2 == 20 {
						R20 = R7 - immediate
					}
					if count2 == 21 {
						R21 = R7 - immediate
					}
					if count2 == 22 {
						R22 = R7 - immediate
					}
					if count2 == 23 {
						R23 = R7 - immediate
					}
					if count2 == 24 {
						R24 = R7 - immediate
					}
					if count2 == 25 {
						R25 = R7 - immediate
					}
					if count2 == 26 {
						R26 = R7 - immediate
					}
					if count2 == 27 {
						R27 = R7 - immediate
					}
					if count2 == 28 {
						R28 = R7 - immediate
					}
					if count2 == 29 {
						R29 = R7 - immediate
					}
					if count2 == 30 {
						R30 = R7 - immediate
					}
					if count2 == 31 {
						R31 = R7 - immediate
					}
				}
				if count1 == 8 {
					if count2 == 0 {
						R0 = R8 - immediate
					}
					if count2 == 1 {
						R1 = R8 - immediate
					}
					if count2 == 2 {
						R2 = R8 - immediate
					}
					if count2 == 3 {
						R3 = R8 - immediate
					}
					if count2 == 4 {
						R4 = R8 - immediate
					}
					if count2 == 5 {
						R5 = R8 - immediate
					}
					if count2 == 6 {
						R6 = R8 - immediate
					}
					if count2 == 7 {
						R7 = R8 - immediate
					}
					if count2 == 8 {
						R8 = R8 - immediate
					}
					if count2 == 9 {
						R9 = R8 - immediate
					}
					if count2 == 10 {
						R10 = R8 - immediate
					}
					if count2 == 11 {
						R11 = R8 - immediate
					}
					if count2 == 12 {
						R12 = R8 - immediate
					}
					if count2 == 13 {
						R13 = R8 - immediate
					}
					if count2 == 14 {
						R14 = R8 - immediate
					}
					if count2 == 15 {
						R15 = R8 - immediate
					}
					if count2 == 16 {
						R16 = R8 - immediate
					}
					if count2 == 17 {
						R17 = R8 - immediate
					}
					if count2 == 18 {
						R18 = R8 - immediate
					}
					if count2 == 19 {
						R19 = R8 - immediate
					}
					if count2 == 20 {
						R20 = R8 - immediate
					}
					if count2 == 21 {
						R21 = R8 - immediate
					}
					if count2 == 22 {
						R22 = R8 - immediate
					}
					if count2 == 23 {
						R23 = R8 - immediate
					}
					if count2 == 24 {
						R24 = R8 - immediate
					}
					if count2 == 25 {
						R25 = R8 - immediate
					}
					if count2 == 26 {
						R26 = R8 - immediate
					}
					if count2 == 27 {
						R27 = R8 - immediate
					}
					if count2 == 28 {
						R28 = R8 - immediate
					}
					if count2 == 29 {
						R29 = R8 - immediate
					}
					if count2 == 30 {
						R30 = R8 - immediate
					}
					if count2 == 31 {
						R31 = R8 - immediate
					}
				}
				if count1 == 9 {
					if count2 == 0 {
						R0 = R9 - immediate
					}
					if count2 == 1 {
						R1 = R9 - immediate
					}
					if count2 == 2 {
						R2 = R9 - immediate
					}
					if count2 == 3 {
						R3 = R9 - immediate
					}
					if count2 == 4 {
						R4 = R9 - immediate
					}
					if count2 == 5 {
						R5 = R9 - immediate
					}
					if count2 == 6 {
						R6 = R9 - immediate
					}
					if count2 == 7 {
						R7 = R9 - immediate
					}
					if count2 == 8 {
						R8 = R9 - immediate
					}
					if count2 == 9 {
						R9 = R9 - immediate
					}
					if count2 == 10 {
						R10 = R9 - immediate
					}
					if count2 == 11 {
						R11 = R9 - immediate
					}
					if count2 == 12 {
						R12 = R9 - immediate
					}
					if count2 == 13 {
						R13 = R9 - immediate
					}
					if count2 == 14 {
						R14 = R9 - immediate
					}
					if count2 == 15 {
						R15 = R9 - immediate
					}
					if count2 == 16 {
						R16 = R9 - immediate
					}
					if count2 == 17 {
						R17 = R9 - immediate
					}
					if count2 == 18 {
						R18 = R9 - immediate
					}
					if count2 == 19 {
						R19 = R9 - immediate
					}
					if count2 == 20 {
						R20 = R9 - immediate
					}
					if count2 == 21 {
						R21 = R9 - immediate
					}
					if count2 == 22 {
						R22 = R9 - immediate
					}
					if count2 == 23 {
						R23 = R9 - immediate
					}
					if count2 == 24 {
						R24 = R9 - immediate
					}
					if count2 == 25 {
						R25 = R9 - immediate
					}
					if count2 == 26 {
						R26 = R9 - immediate
					}
					if count2 == 27 {
						R27 = R9 - immediate
					}
					if count2 == 28 {
						R28 = R9 - immediate
					}
					if count2 == 29 {
						R29 = R9 - immediate
					}
					if count2 == 30 {
						R30 = R9 - immediate
					}
					if count2 == 31 {
						R31 = R9 - immediate
					}
				}
				if count1 == 10 {
					if count2 == 0 {
						R0 = R10 - immediate
					}
					if count2 == 1 {
						R1 = R10 - immediate
					}
					if count2 == 2 {
						R2 = R10 - immediate
					}
					if count2 == 3 {
						R3 = R10 - immediate
					}
					if count2 == 4 {
						R4 = R10 - immediate
					}
					if count2 == 5 {
						R5 = R10 - immediate
					}
					if count2 == 6 {
						R6 = R10 - immediate
					}
					if count2 == 7 {
						R7 = R10 - immediate
					}
					if count2 == 8 {
						R8 = R10 - immediate
					}
					if count2 == 9 {
						R9 = R10 - immediate
					}
					if count2 == 10 {
						R10 = R10 - immediate
					}
					if count2 == 11 {
						R11 = R10 - immediate
					}
					if count2 == 12 {
						R12 = R10 - immediate
					}
					if count2 == 13 {
						R13 = R10 - immediate
					}
					if count2 == 14 {
						R14 = R10 - immediate
					}
					if count2 == 15 {
						R15 = R10 - immediate
					}
					if count2 == 16 {
						R16 = R10 - immediate
					}
					if count2 == 17 {
						R17 = R10 - immediate
					}
					if count2 == 18 {
						R18 = R10 - immediate
					}
					if count2 == 19 {
						R19 = R10 - immediate
					}
					if count2 == 20 {
						R20 = R10 - immediate
					}
					if count2 == 21 {
						R21 = R10 - immediate
					}
					if count2 == 22 {
						R22 = R10 - immediate
					}
					if count2 == 23 {
						R23 = R10 - immediate
					}
					if count2 == 24 {
						R24 = R10 - immediate
					}
					if count2 == 25 {
						R25 = R10 - immediate
					}
					if count2 == 26 {
						R26 = R10 - immediate
					}
					if count2 == 27 {
						R27 = R10 - immediate
					}
					if count2 == 28 {
						R28 = R10 - immediate
					}
					if count2 == 29 {
						R29 = R10 - immediate
					}
					if count2 == 30 {
						R30 = R10 - immediate
					}
					if count2 == 31 {
						R31 = R10 - immediate
					}
				}
				if count1 == 11 {
					if count2 == 0 {
						R0 = R11 - immediate
					}
					if count2 == 1 {
						R1 = R11 - immediate
					}
					if count2 == 2 {
						R2 = R11 - immediate
					}
					if count2 == 3 {
						R3 = R11 - immediate
					}
					if count2 == 4 {
						R4 = R11 - immediate
					}
					if count2 == 5 {
						R5 = R11 - immediate
					}
					if count2 == 6 {
						R6 = R11 - immediate
					}
					if count2 == 7 {
						R7 = R11 - immediate
					}
					if count2 == 8 {
						R8 = R11 - immediate
					}
					if count2 == 9 {
						R9 = R11 - immediate
					}
					if count2 == 10 {
						R10 = R11 - immediate
					}
					if count2 == 11 {
						R11 = R11 - immediate
					}
					if count2 == 12 {
						R12 = R11 - immediate
					}
					if count2 == 13 {
						R13 = R11 - immediate
					}
					if count2 == 14 {
						R14 = R11 - immediate
					}
					if count2 == 15 {
						R15 = R11 - immediate
					}
					if count2 == 16 {
						R16 = R11 - immediate
					}
					if count2 == 17 {
						R17 = R11 - immediate
					}
					if count2 == 18 {
						R18 = R11 - immediate
					}
					if count2 == 19 {
						R19 = R11 - immediate
					}
					if count2 == 20 {
						R20 = R11 - immediate
					}
					if count2 == 21 {
						R21 = R11 - immediate
					}
					if count2 == 22 {
						R22 = R11 - immediate
					}
					if count2 == 23 {
						R23 = R11 - immediate
					}
					if count2 == 24 {
						R24 = R11 - immediate
					}
					if count2 == 25 {
						R25 = R11 - immediate
					}
					if count2 == 26 {
						R26 = R11 - immediate
					}
					if count2 == 27 {
						R27 = R11 - immediate
					}
					if count2 == 28 {
						R28 = R11 - immediate
					}
					if count2 == 29 {
						R29 = R11 - immediate
					}
					if count2 == 30 {
						R30 = R11 - immediate
					}
					if count2 == 31 {
						R31 = R11 - immediate
					}
				}
				if count1 == 12 {
					if count2 == 0 {
						R0 = R12 - immediate
					}
					if count2 == 1 {
						R1 = R12 - immediate
					}
					if count2 == 2 {
						R2 = R12 - immediate
					}
					if count2 == 3 {
						R3 = R12 - immediate
					}
					if count2 == 4 {
						R4 = R12 - immediate
					}
					if count2 == 5 {
						R5 = R12 - immediate
					}
					if count2 == 6 {
						R6 = R12 - immediate
					}
					if count2 == 7 {
						R7 = R12 - immediate
					}
					if count2 == 8 {
						R8 = R12 - immediate
					}
					if count2 == 9 {
						R9 = R12 - immediate
					}
					if count2 == 10 {
						R10 = R12 - immediate
					}
					if count2 == 11 {
						R11 = R12 - immediate
					}
					if count2 == 12 {
						R12 = R12 - immediate
					}
					if count2 == 13 {
						R13 = R12 - immediate
					}
					if count2 == 14 {
						R14 = R12 - immediate
					}
					if count2 == 15 {
						R15 = R12 - immediate
					}
					if count2 == 16 {
						R16 = R12 - immediate
					}
					if count2 == 17 {
						R17 = R12 - immediate
					}
					if count2 == 18 {
						R18 = R12 - immediate
					}
					if count2 == 19 {
						R19 = R12 - immediate
					}
					if count2 == 20 {
						R20 = R12 - immediate
					}
					if count2 == 21 {
						R21 = R12 - immediate
					}
					if count2 == 22 {
						R22 = R12 - immediate
					}
					if count2 == 23 {
						R23 = R12 - immediate
					}
					if count2 == 24 {
						R24 = R12 - immediate
					}
					if count2 == 25 {
						R25 = R12 - immediate
					}
					if count2 == 26 {
						R26 = R12 - immediate
					}
					if count2 == 27 {
						R27 = R12 - immediate
					}
					if count2 == 28 {
						R28 = R12 - immediate
					}
					if count2 == 29 {
						R29 = R12 - immediate
					}
					if count2 == 30 {
						R30 = R12 - immediate
					}
					if count2 == 31 {
						R31 = R12 - immediate
					}
				}
				if count1 == 13 {
					if count2 == 0 {
						R0 = R13 - immediate
					}
					if count2 == 1 {
						R1 = R13 - immediate
					}
					if count2 == 2 {
						R2 = R13 - immediate
					}
					if count2 == 3 {
						R3 = R13 - immediate
					}
					if count2 == 4 {
						R4 = R13 - immediate
					}
					if count2 == 5 {
						R5 = R13 - immediate
					}
					if count2 == 6 {
						R6 = R13 - immediate
					}
					if count2 == 7 {
						R7 = R13 - immediate
					}
					if count2 == 8 {
						R8 = R13 - immediate
					}
					if count2 == 9 {
						R9 = R13 - immediate
					}
					if count2 == 10 {
						R10 = R13 - immediate
					}
					if count2 == 11 {
						R11 = R13 - immediate
					}
					if count2 == 12 {
						R12 = R13 - immediate
					}
					if count2 == 13 {
						R13 = R13 - immediate
					}
					if count2 == 14 {
						R14 = R13 - immediate
					}
					if count2 == 15 {
						R15 = R13 - immediate
					}
					if count2 == 16 {
						R16 = R13 - immediate
					}
					if count2 == 17 {
						R17 = R13 - immediate
					}
					if count2 == 18 {
						R18 = R13 - immediate
					}
					if count2 == 19 {
						R19 = R13 - immediate
					}
					if count2 == 20 {
						R20 = R13 - immediate
					}
					if count2 == 21 {
						R21 = R13 - immediate
					}
					if count2 == 22 {
						R22 = R13 - immediate
					}
					if count2 == 23 {
						R23 = R13 - immediate
					}
					if count2 == 24 {
						R24 = R13 - immediate
					}
					if count2 == 25 {
						R25 = R13 - immediate
					}
					if count2 == 26 {
						R26 = R13 - immediate
					}
					if count2 == 27 {
						R27 = R13 - immediate
					}
					if count2 == 28 {
						R28 = R13 - immediate
					}
					if count2 == 29 {
						R29 = R13 - immediate
					}
					if count2 == 30 {
						R30 = R13 - immediate
					}
					if count2 == 31 {
						R31 = R13 - immediate
					}
				}
				if count1 == 14 {
					if count2 == 0 {
						R0 = R14 - immediate
					}
					if count2 == 1 {
						R1 = R14 - immediate
					}
					if count2 == 2 {
						R2 = R14 - immediate
					}
					if count2 == 3 {
						R3 = R14 - immediate
					}
					if count2 == 4 {
						R4 = R14 - immediate
					}
					if count2 == 5 {
						R5 = R14 - immediate
					}
					if count2 == 6 {
						R6 = R14 - immediate
					}
					if count2 == 7 {
						R7 = R14 - immediate
					}
					if count2 == 8 {
						R8 = R14 - immediate
					}
					if count2 == 9 {
						R9 = R14 - immediate
					}
					if count2 == 10 {
						R10 = R14 - immediate
					}
					if count2 == 11 {
						R11 = R14 - immediate
					}
					if count2 == 12 {
						R12 = R14 - immediate
					}
					if count2 == 13 {
						R13 = R14 - immediate
					}
					if count2 == 14 {
						R14 = R14 - immediate
					}
					if count2 == 15 {
						R15 = R14 - immediate
					}
					if count2 == 16 {
						R16 = R14 - immediate
					}
					if count2 == 17 {
						R17 = R14 - immediate
					}
					if count2 == 18 {
						R18 = R14 - immediate
					}
					if count2 == 19 {
						R19 = R14 - immediate
					}
					if count2 == 20 {
						R20 = R14 - immediate
					}
					if count2 == 21 {
						R21 = R14 - immediate
					}
					if count2 == 22 {
						R22 = R14 - immediate
					}
					if count2 == 23 {
						R23 = R14 - immediate
					}
					if count2 == 24 {
						R24 = R14 - immediate
					}
					if count2 == 25 {
						R25 = R14 - immediate
					}
					if count2 == 26 {
						R26 = R14 - immediate
					}
					if count2 == 27 {
						R27 = R14 - immediate
					}
					if count2 == 28 {
						R28 = R14 - immediate
					}
					if count2 == 29 {
						R29 = R14 - immediate
					}
					if count2 == 30 {
						R30 = R14 - immediate
					}
					if count2 == 31 {
						R31 = R14 - immediate
					}
				}
				if count1 == 15 {
					if count2 == 0 {
						R0 = R15 - immediate
					}
					if count2 == 1 {
						R1 = R15 - immediate
					}
					if count2 == 2 {
						R2 = R15 - immediate
					}
					if count2 == 3 {
						R3 = R15 - immediate
					}
					if count2 == 4 {
						R4 = R15 - immediate
					}
					if count2 == 5 {
						R5 = R15 - immediate
					}
					if count2 == 6 {
						R6 = R15 - immediate
					}
					if count2 == 7 {
						R7 = R15 - immediate
					}
					if count2 == 8 {
						R8 = R15 - immediate
					}
					if count2 == 9 {
						R9 = R15 - immediate
					}
					if count2 == 10 {
						R10 = R15 - immediate
					}
					if count2 == 11 {
						R11 = R15 - immediate
					}
					if count2 == 12 {
						R12 = R15 - immediate
					}
					if count2 == 13 {
						R13 = R15 - immediate
					}
					if count2 == 14 {
						R14 = R15 - immediate
					}
					if count2 == 15 {
						R15 = R15 - immediate
					}
					if count2 == 16 {
						R16 = R15 - immediate
					}
					if count2 == 17 {
						R17 = R15 - immediate
					}
					if count2 == 18 {
						R18 = R15 - immediate
					}
					if count2 == 19 {
						R19 = R15 - immediate
					}
					if count2 == 20 {
						R20 = R15 - immediate
					}
					if count2 == 21 {
						R21 = R15 - immediate
					}
					if count2 == 22 {
						R22 = R15 - immediate
					}
					if count2 == 23 {
						R23 = R15 - immediate
					}
					if count2 == 24 {
						R24 = R15 - immediate
					}
					if count2 == 25 {
						R25 = R15 - immediate
					}
					if count2 == 26 {
						R26 = R15 - immediate
					}
					if count2 == 27 {
						R27 = R15 - immediate
					}
					if count2 == 28 {
						R28 = R15 - immediate
					}
					if count2 == 29 {
						R29 = R15 - immediate
					}
					if count2 == 30 {
						R30 = R15 - immediate
					}
					if count2 == 31 {
						R31 = R15 - immediate
					}
				}
				if count1 == 16 {
					if count2 == 0 {
						R0 = R16 - immediate
					}
					if count2 == 1 {
						R1 = R16 - immediate
					}
					if count2 == 2 {
						R2 = R16 - immediate
					}
					if count2 == 3 {
						R3 = R16 - immediate
					}
					if count2 == 4 {
						R4 = R16 - immediate
					}
					if count2 == 5 {
						R5 = R16 - immediate
					}
					if count2 == 6 {
						R6 = R16 - immediate
					}
					if count2 == 7 {
						R7 = R16 - immediate
					}
					if count2 == 8 {
						R8 = R16 - immediate
					}
					if count2 == 9 {
						R9 = R16 - immediate
					}
					if count2 == 10 {
						R10 = R16 - immediate
					}
					if count2 == 11 {
						R11 = R16 - immediate
					}
					if count2 == 12 {
						R12 = R16 - immediate
					}
					if count2 == 13 {
						R13 = R16 - immediate
					}
					if count2 == 14 {
						R14 = R16 - immediate
					}
					if count2 == 15 {
						R15 = R16 - immediate
					}
					if count2 == 16 {
						R16 = R16 - immediate
					}
					if count2 == 17 {
						R17 = R16 - immediate
					}
					if count2 == 18 {
						R18 = R16 - immediate
					}
					if count2 == 19 {
						R19 = R16 - immediate
					}
					if count2 == 20 {
						R20 = R16 - immediate
					}
					if count2 == 21 {
						R21 = R16 - immediate
					}
					if count2 == 22 {
						R22 = R16 - immediate
					}
					if count2 == 23 {
						R23 = R16 - immediate
					}
					if count2 == 24 {
						R24 = R16 - immediate
					}
					if count2 == 25 {
						R25 = R16 - immediate
					}
					if count2 == 26 {
						R26 = R16 - immediate
					}
					if count2 == 27 {
						R27 = R16 - immediate
					}
					if count2 == 28 {
						R28 = R16 - immediate
					}
					if count2 == 29 {
						R29 = R16 - immediate
					}
					if count2 == 30 {
						R30 = R16 - immediate
					}
					if count2 == 31 {
						R31 = R16 - immediate
					}
				}
				if count1 == 17 {
					if count2 == 0 {
						R0 = R17 - immediate
					}
					if count2 == 1 {
						R1 = R17 - immediate
					}
					if count2 == 2 {
						R2 = R17 - immediate
					}
					if count2 == 3 {
						R3 = R17 - immediate
					}
					if count2 == 4 {
						R4 = R17 - immediate
					}
					if count2 == 5 {
						R5 = R17 - immediate
					}
					if count2 == 6 {
						R6 = R17 - immediate
					}
					if count2 == 7 {
						R7 = R17 - immediate
					}
					if count2 == 8 {
						R8 = R17 - immediate
					}
					if count2 == 9 {
						R9 = R17 - immediate
					}
					if count2 == 10 {
						R10 = R17 - immediate
					}
					if count2 == 11 {
						R11 = R17 - immediate
					}
					if count2 == 12 {
						R12 = R17 - immediate
					}
					if count2 == 13 {
						R13 = R17 - immediate
					}
					if count2 == 14 {
						R14 = R17 - immediate
					}
					if count2 == 15 {
						R15 = R17 - immediate
					}
					if count2 == 16 {
						R16 = R17 - immediate
					}
					if count2 == 17 {
						R17 = R17 - immediate
					}
					if count2 == 18 {
						R18 = R17 - immediate
					}
					if count2 == 19 {
						R19 = R17 - immediate
					}
					if count2 == 20 {
						R20 = R17 - immediate
					}
					if count2 == 21 {
						R21 = R17 - immediate
					}
					if count2 == 22 {
						R22 = R17 - immediate
					}
					if count2 == 23 {
						R23 = R17 - immediate
					}
					if count2 == 24 {
						R24 = R17 - immediate
					}
					if count2 == 25 {
						R25 = R17 - immediate
					}
					if count2 == 26 {
						R26 = R17 - immediate
					}
					if count2 == 27 {
						R27 = R17 - immediate
					}
					if count2 == 28 {
						R28 = R17 - immediate
					}
					if count2 == 29 {
						R29 = R17 - immediate
					}
					if count2 == 30 {
						R30 = R17 - immediate
					}
					if count2 == 31 {
						R31 = R17 - immediate
					}
				}
				if count1 == 18 {
					if count2 == 0 {
						R0 = R18 - immediate
					}
					if count2 == 1 {
						R1 = R18 - immediate
					}
					if count2 == 2 {
						R2 = R18 - immediate
					}
					if count2 == 3 {
						R3 = R18 - immediate
					}
					if count2 == 4 {
						R4 = R18 - immediate
					}
					if count2 == 5 {
						R5 = R18 - immediate
					}
					if count2 == 6 {
						R6 = R18 - immediate
					}
					if count2 == 7 {
						R7 = R18 - immediate
					}
					if count2 == 8 {
						R8 = R18 - immediate
					}
					if count2 == 9 {
						R9 = R18 - immediate
					}
					if count2 == 10 {
						R10 = R18 - immediate
					}
					if count2 == 11 {
						R11 = R18 - immediate
					}
					if count2 == 12 {
						R12 = R18 - immediate
					}
					if count2 == 13 {
						R13 = R18 - immediate
					}
					if count2 == 14 {
						R14 = R18 - immediate
					}
					if count2 == 15 {
						R15 = R18 - immediate
					}
					if count2 == 16 {
						R16 = R18 - immediate
					}
					if count2 == 17 {
						R17 = R18 - immediate
					}
					if count2 == 18 {
						R18 = R18 - immediate
					}
					if count2 == 19 {
						R19 = R18 - immediate
					}
					if count2 == 20 {
						R20 = R18 - immediate
					}
					if count2 == 21 {
						R21 = R18 - immediate
					}
					if count2 == 22 {
						R22 = R18 - immediate
					}
					if count2 == 23 {
						R23 = R18 - immediate
					}
					if count2 == 24 {
						R24 = R18 - immediate
					}
					if count2 == 25 {
						R25 = R18 - immediate
					}
					if count2 == 26 {
						R26 = R18 - immediate
					}
					if count2 == 27 {
						R27 = R18 - immediate
					}
					if count2 == 28 {
						R28 = R18 - immediate
					}
					if count2 == 29 {
						R29 = R18 - immediate
					}
					if count2 == 30 {
						R30 = R18 - immediate
					}
					if count2 == 31 {
						R31 = R18 - immediate
					}
				}
				if count1 == 19 {
					if count2 == 0 {
						R0 = R19 - immediate
					}
					if count2 == 1 {
						R1 = R19 - immediate
					}
					if count2 == 2 {
						R2 = R19 - immediate
					}
					if count2 == 3 {
						R3 = R19 - immediate
					}
					if count2 == 4 {
						R4 = R19 - immediate
					}
					if count2 == 5 {
						R5 = R19 - immediate
					}
					if count2 == 6 {
						R6 = R19 - immediate
					}
					if count2 == 7 {
						R7 = R19 - immediate
					}
					if count2 == 8 {
						R8 = R19 - immediate
					}
					if count2 == 9 {
						R9 = R19 - immediate
					}
					if count2 == 10 {
						R10 = R19 - immediate
					}
					if count2 == 11 {
						R11 = R19 - immediate
					}
					if count2 == 12 {
						R12 = R19 - immediate
					}
					if count2 == 13 {
						R13 = R19 - immediate
					}
					if count2 == 14 {
						R14 = R19 - immediate
					}
					if count2 == 15 {
						R15 = R19 - immediate
					}
					if count2 == 16 {
						R16 = R19 - immediate
					}
					if count2 == 17 {
						R17 = R19 - immediate
					}
					if count2 == 18 {
						R18 = R19 - immediate
					}
					if count2 == 19 {
						R19 = R19 - immediate
					}
					if count2 == 20 {
						R20 = R19 - immediate
					}
					if count2 == 21 {
						R21 = R19 - immediate
					}
					if count2 == 22 {
						R22 = R19 - immediate
					}
					if count2 == 23 {
						R23 = R19 - immediate
					}
					if count2 == 24 {
						R24 = R19 - immediate
					}
					if count2 == 25 {
						R25 = R19 - immediate
					}
					if count2 == 26 {
						R26 = R19 - immediate
					}
					if count2 == 27 {
						R27 = R19 - immediate
					}
					if count2 == 28 {
						R28 = R19 - immediate
					}
					if count2 == 29 {
						R29 = R19 - immediate
					}
					if count2 == 30 {
						R30 = R19 - immediate
					}
					if count2 == 31 {
						R31 = R19 - immediate
					}
				}
				if count1 == 20 {
					if count2 == 0 {
						R0 = R20 - immediate
					}
					if count2 == 1 {
						R1 = R20 - immediate
					}
					if count2 == 2 {
						R2 = R20 - immediate
					}
					if count2 == 3 {
						R3 = R20 - immediate
					}
					if count2 == 4 {
						R4 = R20 - immediate
					}
					if count2 == 5 {
						R5 = R20 - immediate
					}
					if count2 == 6 {
						R6 = R20 - immediate
					}
					if count2 == 7 {
						R7 = R20 - immediate
					}
					if count2 == 8 {
						R8 = R20 - immediate
					}
					if count2 == 9 {
						R9 = R20 - immediate
					}
					if count2 == 10 {
						R10 = R20 - immediate
					}
					if count2 == 11 {
						R11 = R20 - immediate
					}
					if count2 == 12 {
						R12 = R20 - immediate
					}
					if count2 == 13 {
						R13 = R20 - immediate
					}
					if count2 == 14 {
						R14 = R20 - immediate
					}
					if count2 == 15 {
						R15 = R20 - immediate
					}
					if count2 == 16 {
						R16 = R20 - immediate
					}
					if count2 == 17 {
						R17 = R20 - immediate
					}
					if count2 == 18 {
						R18 = R20 - immediate
					}
					if count2 == 19 {
						R19 = R20 - immediate
					}
					if count2 == 20 {
						R20 = R20 - immediate
					}
					if count2 == 21 {
						R21 = R20 - immediate
					}
					if count2 == 22 {
						R22 = R20 - immediate
					}
					if count2 == 23 {
						R23 = R20 - immediate
					}
					if count2 == 24 {
						R24 = R20 - immediate
					}
					if count2 == 25 {
						R25 = R20 - immediate
					}
					if count2 == 26 {
						R26 = R20 - immediate
					}
					if count2 == 27 {
						R27 = R20 - immediate
					}
					if count2 == 28 {
						R28 = R20 - immediate
					}
					if count2 == 29 {
						R29 = R20 - immediate
					}
					if count2 == 30 {
						R30 = R20 - immediate
					}
					if count2 == 31 {
						R31 = R20 - immediate
					}
				}
				if count1 == 21 {
					if count2 == 0 {
						R0 = R21 - immediate
					}
					if count2 == 1 {
						R1 = R21 - immediate
					}
					if count2 == 2 {
						R2 = R21 - immediate
					}
					if count2 == 3 {
						R3 = R21 - immediate
					}
					if count2 == 4 {
						R4 = R21 - immediate
					}
					if count2 == 5 {
						R5 = R21 - immediate
					}
					if count2 == 6 {
						R6 = R21 - immediate
					}
					if count2 == 7 {
						R7 = R21 - immediate
					}
					if count2 == 8 {
						R8 = R21 - immediate
					}
					if count2 == 9 {
						R9 = R21 - immediate
					}
					if count2 == 10 {
						R10 = R21 - immediate
					}
					if count2 == 11 {
						R11 = R21 - immediate
					}
					if count2 == 12 {
						R12 = R21 - immediate
					}
					if count2 == 13 {
						R13 = R21 - immediate
					}
					if count2 == 14 {
						R14 = R21 - immediate
					}
					if count2 == 15 {
						R15 = R21 - immediate
					}
					if count2 == 16 {
						R16 = R21 - immediate
					}
					if count2 == 17 {
						R17 = R21 - immediate
					}
					if count2 == 18 {
						R18 = R21 - immediate
					}
					if count2 == 19 {
						R19 = R21 - immediate
					}
					if count2 == 20 {
						R20 = R21 - immediate
					}
					if count2 == 21 {
						R21 = R21 - immediate
					}
					if count2 == 22 {
						R22 = R21 - immediate
					}
					if count2 == 23 {
						R23 = R21 - immediate
					}
					if count2 == 24 {
						R24 = R21 - immediate
					}
					if count2 == 25 {
						R25 = R21 - immediate
					}
					if count2 == 26 {
						R26 = R21 - immediate
					}
					if count2 == 27 {
						R27 = R21 - immediate
					}
					if count2 == 28 {
						R28 = R21 - immediate
					}
					if count2 == 29 {
						R29 = R21 - immediate
					}
					if count2 == 30 {
						R30 = R21 - immediate
					}
					if count2 == 31 {
						R31 = R21 - immediate
					}
				}
				if count1 == 22 {
					if count2 == 0 {
						R0 = R22 - immediate
					}
					if count2 == 1 {
						R1 = R22 - immediate
					}
					if count2 == 2 {
						R2 = R22 - immediate
					}
					if count2 == 3 {
						R3 = R22 - immediate
					}
					if count2 == 4 {
						R4 = R22 - immediate
					}
					if count2 == 5 {
						R5 = R22 - immediate
					}
					if count2 == 6 {
						R6 = R22 - immediate
					}
					if count2 == 7 {
						R7 = R22 - immediate
					}
					if count2 == 8 {
						R8 = R22 - immediate
					}
					if count2 == 9 {
						R9 = R22 - immediate
					}
					if count2 == 10 {
						R10 = R22 - immediate
					}
					if count2 == 11 {
						R11 = R22 - immediate
					}
					if count2 == 12 {
						R12 = R22 - immediate
					}
					if count2 == 13 {
						R13 = R22 - immediate
					}
					if count2 == 14 {
						R14 = R22 - immediate
					}
					if count2 == 15 {
						R15 = R22 - immediate
					}
					if count2 == 16 {
						R16 = R22 - immediate
					}
					if count2 == 17 {
						R17 = R22 - immediate
					}
					if count2 == 18 {
						R18 = R22 - immediate
					}
					if count2 == 19 {
						R19 = R22 - immediate
					}
					if count2 == 20 {
						R20 = R22 - immediate
					}
					if count2 == 21 {
						R21 = R22 - immediate
					}
					if count2 == 22 {
						R22 = R22 - immediate
					}
					if count2 == 23 {
						R23 = R22 - immediate
					}
					if count2 == 24 {
						R24 = R22 - immediate
					}
					if count2 == 25 {
						R25 = R22 - immediate
					}
					if count2 == 26 {
						R26 = R22 - immediate
					}
					if count2 == 27 {
						R27 = R22 - immediate
					}
					if count2 == 28 {
						R28 = R22 - immediate
					}
					if count2 == 29 {
						R29 = R22 - immediate
					}
					if count2 == 30 {
						R30 = R22 - immediate
					}
					if count2 == 31 {
						R31 = R22 - immediate
					}
				}
				if count1 == 23 {
					if count2 == 0 {
						R0 = R23 - immediate
					}
					if count2 == 1 {
						R1 = R23 - immediate
					}
					if count2 == 2 {
						R2 = R23 - immediate
					}
					if count2 == 3 {
						R3 = R23 - immediate
					}
					if count2 == 4 {
						R4 = R23 - immediate
					}
					if count2 == 5 {
						R5 = R23 - immediate
					}
					if count2 == 6 {
						R6 = R23 - immediate
					}
					if count2 == 7 {
						R7 = R23 - immediate
					}
					if count2 == 8 {
						R8 = R23 - immediate
					}
					if count2 == 9 {
						R9 = R23 - immediate
					}
					if count2 == 10 {
						R10 = R23 - immediate
					}
					if count2 == 11 {
						R11 = R23 - immediate
					}
					if count2 == 12 {
						R12 = R23 - immediate
					}
					if count2 == 13 {
						R13 = R23 - immediate
					}
					if count2 == 14 {
						R14 = R23 - immediate
					}
					if count2 == 15 {
						R15 = R23 - immediate
					}
					if count2 == 16 {
						R16 = R23 - immediate
					}
					if count2 == 17 {
						R17 = R23 - immediate
					}
					if count2 == 18 {
						R18 = R23 - immediate
					}
					if count2 == 19 {
						R19 = R23 - immediate
					}
					if count2 == 20 {
						R20 = R23 - immediate
					}
					if count2 == 21 {
						R21 = R23 - immediate
					}
					if count2 == 22 {
						R22 = R23 - immediate
					}
					if count2 == 23 {
						R23 = R23 - immediate
					}
					if count2 == 24 {
						R24 = R23 - immediate
					}
					if count2 == 25 {
						R25 = R23 - immediate
					}
					if count2 == 26 {
						R26 = R23 - immediate
					}
					if count2 == 27 {
						R27 = R23 - immediate
					}
					if count2 == 28 {
						R28 = R23 - immediate
					}
					if count2 == 29 {
						R29 = R23 - immediate
					}
					if count2 == 30 {
						R30 = R23 - immediate
					}
					if count2 == 31 {
						R31 = R23 - immediate
					}
				}
				if count1 == 24 {
					if count2 == 0 {
						R0 = R24 - immediate
					}
					if count2 == 1 {
						R1 = R24 - immediate
					}
					if count2 == 2 {
						R2 = R24 - immediate
					}
					if count2 == 3 {
						R3 = R24 - immediate
					}
					if count2 == 4 {
						R4 = R24 - immediate
					}
					if count2 == 5 {
						R5 = R24 - immediate
					}
					if count2 == 6 {
						R6 = R24 - immediate
					}
					if count2 == 7 {
						R7 = R24 - immediate
					}
					if count2 == 8 {
						R8 = R24 - immediate
					}
					if count2 == 9 {
						R9 = R24 - immediate
					}
					if count2 == 10 {
						R10 = R24 - immediate
					}
					if count2 == 11 {
						R11 = R24 - immediate
					}
					if count2 == 12 {
						R12 = R24 - immediate
					}
					if count2 == 13 {
						R13 = R24 - immediate
					}
					if count2 == 14 {
						R14 = R24 - immediate
					}
					if count2 == 15 {
						R15 = R24 - immediate
					}
					if count2 == 16 {
						R16 = R24 - immediate
					}
					if count2 == 17 {
						R17 = R24 - immediate
					}
					if count2 == 18 {
						R18 = R24 - immediate
					}
					if count2 == 19 {
						R19 = R24 - immediate
					}
					if count2 == 20 {
						R20 = R24 - immediate
					}
					if count2 == 21 {
						R21 = R24 - immediate
					}
					if count2 == 22 {
						R22 = R24 - immediate
					}
					if count2 == 23 {
						R23 = R24 - immediate
					}
					if count2 == 24 {
						R24 = R24 - immediate
					}
					if count2 == 25 {
						R25 = R24 - immediate
					}
					if count2 == 26 {
						R26 = R24 - immediate
					}
					if count2 == 27 {
						R27 = R24 - immediate
					}
					if count2 == 28 {
						R28 = R24 - immediate
					}
					if count2 == 29 {
						R29 = R24 - immediate
					}
					if count2 == 30 {
						R30 = R24 - immediate
					}
					if count2 == 31 {
						R31 = R24 - immediate
					}
				}
				if count1 == 25 {
					if count2 == 0 {
						R0 = R25 - immediate
					}
					if count2 == 1 {
						R1 = R25 - immediate
					}
					if count2 == 2 {
						R2 = R25 - immediate
					}
					if count2 == 3 {
						R3 = R25 - immediate
					}
					if count2 == 4 {
						R4 = R25 - immediate
					}
					if count2 == 5 {
						R5 = R25 - immediate
					}
					if count2 == 6 {
						R6 = R25 - immediate
					}
					if count2 == 7 {
						R7 = R25 - immediate
					}
					if count2 == 8 {
						R8 = R25 - immediate
					}
					if count2 == 9 {
						R9 = R25 - immediate
					}
					if count2 == 10 {
						R10 = R25 - immediate
					}
					if count2 == 11 {
						R11 = R25 - immediate
					}
					if count2 == 12 {
						R12 = R25 - immediate
					}
					if count2 == 13 {
						R13 = R25 - immediate
					}
					if count2 == 14 {
						R14 = R25 - immediate
					}
					if count2 == 15 {
						R15 = R25 - immediate
					}
					if count2 == 16 {
						R16 = R25 - immediate
					}
					if count2 == 17 {
						R17 = R25 - immediate
					}
					if count2 == 18 {
						R18 = R25 - immediate
					}
					if count2 == 19 {
						R19 = R25 - immediate
					}
					if count2 == 20 {
						R20 = R25 - immediate
					}
					if count2 == 21 {
						R21 = R25 - immediate
					}
					if count2 == 22 {
						R22 = R25 - immediate
					}
					if count2 == 23 {
						R23 = R25 - immediate
					}
					if count2 == 24 {
						R24 = R25 - immediate
					}
					if count2 == 25 {
						R25 = R25 - immediate
					}
					if count2 == 26 {
						R26 = R25 - immediate
					}
					if count2 == 27 {
						R27 = R25 - immediate
					}
					if count2 == 28 {
						R28 = R25 - immediate
					}
					if count2 == 29 {
						R29 = R25 - immediate
					}
					if count2 == 30 {
						R30 = R25 - immediate
					}
					if count2 == 31 {
						R31 = R25 - immediate
					}
				}
				if count1 == 26 {
					if count2 == 0 {
						R0 = R26 - immediate
					}
					if count2 == 1 {
						R1 = R26 - immediate
					}
					if count2 == 2 {
						R2 = R26 - immediate
					}
					if count2 == 3 {
						R3 = R26 - immediate
					}
					if count2 == 4 {
						R4 = R26 - immediate
					}
					if count2 == 5 {
						R5 = R26 - immediate
					}
					if count2 == 6 {
						R6 = R26 - immediate
					}
					if count2 == 7 {
						R7 = R26 - immediate
					}
					if count2 == 8 {
						R8 = R26 - immediate
					}
					if count2 == 9 {
						R9 = R26 - immediate
					}
					if count2 == 10 {
						R10 = R26 - immediate
					}
					if count2 == 11 {
						R11 = R26 - immediate
					}
					if count2 == 12 {
						R12 = R26 - immediate
					}
					if count2 == 13 {
						R13 = R26 - immediate
					}
					if count2 == 14 {
						R14 = R26 - immediate
					}
					if count2 == 15 {
						R15 = R26 - immediate
					}
					if count2 == 16 {
						R16 = R26 - immediate
					}
					if count2 == 17 {
						R17 = R26 - immediate
					}
					if count2 == 18 {
						R18 = R26 - immediate
					}
					if count2 == 19 {
						R19 = R26 - immediate
					}
					if count2 == 20 {
						R20 = R26 - immediate
					}
					if count2 == 21 {
						R21 = R26 - immediate
					}
					if count2 == 22 {
						R22 = R26 - immediate
					}
					if count2 == 23 {
						R23 = R26 - immediate
					}
					if count2 == 24 {
						R24 = R26 - immediate
					}
					if count2 == 25 {
						R25 = R26 - immediate
					}
					if count2 == 26 {
						R26 = R26 - immediate
					}
					if count2 == 27 {
						R27 = R26 - immediate
					}
					if count2 == 28 {
						R28 = R26 - immediate
					}
					if count2 == 29 {
						R29 = R26 - immediate
					}
					if count2 == 30 {
						R30 = R26 - immediate
					}
					if count2 == 31 {
						R31 = R26 - immediate
					}
				}
				if count1 == 27 {
					if count2 == 0 {
						R0 = R27 - immediate
					}
					if count2 == 1 {
						R1 = R27 - immediate
					}
					if count2 == 2 {
						R2 = R27 - immediate
					}
					if count2 == 3 {
						R3 = R27 - immediate
					}
					if count2 == 4 {
						R4 = R27 - immediate
					}
					if count2 == 5 {
						R5 = R27 - immediate
					}
					if count2 == 6 {
						R6 = R27 - immediate
					}
					if count2 == 7 {
						R7 = R27 - immediate
					}
					if count2 == 8 {
						R8 = R27 - immediate
					}
					if count2 == 9 {
						R9 = R27 - immediate
					}
					if count2 == 10 {
						R10 = R27 - immediate
					}
					if count2 == 11 {
						R11 = R27 - immediate
					}
					if count2 == 12 {
						R12 = R27 - immediate
					}
					if count2 == 13 {
						R13 = R27 - immediate
					}
					if count2 == 14 {
						R14 = R27 - immediate
					}
					if count2 == 15 {
						R15 = R27 - immediate
					}
					if count2 == 16 {
						R16 = R27 - immediate
					}
					if count2 == 17 {
						R17 = R27 - immediate
					}
					if count2 == 18 {
						R18 = R27 - immediate
					}
					if count2 == 19 {
						R19 = R27 - immediate
					}
					if count2 == 20 {
						R20 = R27 - immediate
					}
					if count2 == 21 {
						R21 = R27 - immediate
					}
					if count2 == 22 {
						R22 = R27 - immediate
					}
					if count2 == 23 {
						R23 = R27 - immediate
					}
					if count2 == 24 {
						R24 = R27 - immediate
					}
					if count2 == 25 {
						R25 = R27 - immediate
					}
					if count2 == 26 {
						R26 = R27 - immediate
					}
					if count2 == 27 {
						R27 = R27 - immediate
					}
					if count2 == 28 {
						R28 = R27 - immediate
					}
					if count2 == 29 {
						R29 = R27 - immediate
					}
					if count2 == 30 {
						R30 = R27 - immediate
					}
					if count2 == 31 {
						R31 = R27 - immediate
					}
				}
				if count1 == 28 {
					if count2 == 0 {
						R0 = R28 - immediate
					}
					if count2 == 1 {
						R1 = R28 - immediate
					}
					if count2 == 2 {
						R2 = R28 - immediate
					}
					if count2 == 3 {
						R3 = R28 - immediate
					}
					if count2 == 4 {
						R4 = R28 - immediate
					}
					if count2 == 5 {
						R5 = R28 - immediate
					}
					if count2 == 6 {
						R6 = R28 - immediate
					}
					if count2 == 7 {
						R7 = R28 - immediate
					}
					if count2 == 8 {
						R8 = R28 - immediate
					}
					if count2 == 9 {
						R9 = R28 - immediate
					}
					if count2 == 10 {
						R10 = R28 - immediate
					}
					if count2 == 11 {
						R11 = R28 - immediate
					}
					if count2 == 12 {
						R12 = R28 - immediate
					}
					if count2 == 13 {
						R13 = R28 - immediate
					}
					if count2 == 14 {
						R14 = R28 - immediate
					}
					if count2 == 15 {
						R15 = R28 - immediate
					}
					if count2 == 16 {
						R16 = R28 - immediate
					}
					if count2 == 17 {
						R17 = R28 - immediate
					}
					if count2 == 18 {
						R18 = R28 - immediate
					}
					if count2 == 19 {
						R19 = R28 - immediate
					}
					if count2 == 20 {
						R20 = R28 - immediate
					}
					if count2 == 21 {
						R21 = R28 - immediate
					}
					if count2 == 22 {
						R22 = R28 - immediate
					}
					if count2 == 23 {
						R23 = R28 - immediate
					}
					if count2 == 24 {
						R24 = R28 - immediate
					}
					if count2 == 25 {
						R25 = R28 - immediate
					}
					if count2 == 26 {
						R26 = R28 - immediate
					}
					if count2 == 27 {
						R27 = R28 - immediate
					}
					if count2 == 28 {
						R28 = R28 - immediate
					}
					if count2 == 29 {
						R29 = R28 - immediate
					}
					if count2 == 30 {
						R30 = R28 - immediate
					}
					if count2 == 31 {
						R31 = R28 - immediate
					}
				}
				if count1 == 29 {
					if count2 == 0 {
						R0 = R29 - immediate
					}
					if count2 == 1 {
						R1 = R29 - immediate
					}
					if count2 == 2 {
						R2 = R29 - immediate
					}
					if count2 == 3 {
						R3 = R29 - immediate
					}
					if count2 == 4 {
						R4 = R29 - immediate
					}
					if count2 == 5 {
						R5 = R29 - immediate
					}
					if count2 == 6 {
						R6 = R29 - immediate
					}
					if count2 == 7 {
						R7 = R29 - immediate
					}
					if count2 == 8 {
						R8 = R29 - immediate
					}
					if count2 == 9 {
						R9 = R29 - immediate
					}
					if count2 == 10 {
						R10 = R29 - immediate
					}
					if count2 == 11 {
						R11 = R29 - immediate
					}
					if count2 == 12 {
						R12 = R29 - immediate
					}
					if count2 == 13 {
						R13 = R29 - immediate
					}
					if count2 == 14 {
						R14 = R29 - immediate
					}
					if count2 == 15 {
						R15 = R29 - immediate
					}
					if count2 == 16 {
						R16 = R29 - immediate
					}
					if count2 == 17 {
						R17 = R29 - immediate
					}
					if count2 == 18 {
						R18 = R29 - immediate
					}
					if count2 == 19 {
						R19 = R29 - immediate
					}
					if count2 == 20 {
						R20 = R29 - immediate
					}
					if count2 == 21 {
						R21 = R29 - immediate
					}
					if count2 == 22 {
						R22 = R29 - immediate
					}
					if count2 == 23 {
						R23 = R29 - immediate
					}
					if count2 == 24 {
						R24 = R29 - immediate
					}
					if count2 == 25 {
						R25 = R29 - immediate
					}
					if count2 == 26 {
						R26 = R29 - immediate
					}
					if count2 == 27 {
						R27 = R29 - immediate
					}
					if count2 == 28 {
						R28 = R29 - immediate
					}
					if count2 == 29 {
						R29 = R29 - immediate
					}
					if count2 == 30 {
						R30 = R29 - immediate
					}
					if count2 == 31 {
						R31 = R29 - immediate
					}
				}
				if count1 == 30 {
					if count2 == 0 {
						R0 = R30 - immediate
					}
					if count2 == 1 {
						R1 = R30 - immediate
					}
					if count2 == 2 {
						R2 = R30 - immediate
					}
					if count2 == 3 {
						R3 = R30 - immediate
					}
					if count2 == 4 {
						R4 = R30 - immediate
					}
					if count2 == 5 {
						R5 = R30 - immediate
					}
					if count2 == 6 {
						R6 = R30 - immediate
					}
					if count2 == 7 {
						R7 = R30 - immediate
					}
					if count2 == 8 {
						R8 = R30 - immediate
					}
					if count2 == 9 {
						R9 = R30 - immediate
					}
					if count2 == 10 {
						R10 = R30 - immediate
					}
					if count2 == 11 {
						R11 = R30 - immediate
					}
					if count2 == 12 {
						R12 = R30 - immediate
					}
					if count2 == 13 {
						R13 = R30 - immediate
					}
					if count2 == 14 {
						R14 = R30 - immediate
					}
					if count2 == 15 {
						R15 = R30 - immediate
					}
					if count2 == 16 {
						R16 = R30 - immediate
					}
					if count2 == 17 {
						R17 = R30 - immediate
					}
					if count2 == 18 {
						R18 = R30 - immediate
					}
					if count2 == 19 {
						R19 = R30 - immediate
					}
					if count2 == 20 {
						R20 = R30 - immediate
					}
					if count2 == 21 {
						R21 = R30 - immediate
					}
					if count2 == 22 {
						R22 = R30 - immediate
					}
					if count2 == 23 {
						R23 = R30 - immediate
					}
					if count2 == 24 {
						R24 = R30 - immediate
					}
					if count2 == 25 {
						R25 = R30 - immediate
					}
					if count2 == 26 {
						R26 = R30 - immediate
					}
					if count2 == 27 {
						R27 = R30 - immediate
					}
					if count2 == 28 {
						R28 = R30 - immediate
					}
					if count2 == 29 {
						R29 = R30 - immediate
					}
					if count2 == 30 {
						R30 = R30 - immediate
					}
					if count2 == 31 {
						R31 = R30 - immediate
					}
				}
				if count1 == 31 {
					if count2 == 0 {
						R0 = R31 - immediate
					}
					if count2 == 1 {
						R1 = R31 - immediate
					}
					if count2 == 2 {
						R2 = R31 - immediate
					}
					if count2 == 3 {
						R3 = R31 - immediate
					}
					if count2 == 4 {
						R4 = R31 - immediate
					}
					if count2 == 5 {
						R5 = R31 - immediate
					}
					if count2 == 6 {
						R6 = R31 - immediate
					}
					if count2 == 7 {
						R7 = R31 - immediate
					}
					if count2 == 8 {
						R8 = R31 - immediate
					}
					if count2 == 9 {
						R9 = R31 - immediate
					}
					if count2 == 10 {
						R10 = R31 - immediate
					}
					if count2 == 11 {
						R11 = R31 - immediate
					}
					if count2 == 12 {
						R12 = R31 - immediate
					}
					if count2 == 13 {
						R13 = R31 - immediate
					}
					if count2 == 14 {
						R14 = R31 - immediate
					}
					if count2 == 15 {
						R15 = R31 - immediate
					}
					if count2 == 16 {
						R16 = R31 - immediate
					}
					if count2 == 17 {
						R17 = R31 - immediate
					}
					if count2 == 18 {
						R18 = R31 - immediate
					}
					if count2 == 19 {
						R19 = R31 - immediate
					}
					if count2 == 20 {
						R20 = R31 - immediate
					}
					if count2 == 21 {
						R21 = R31 - immediate
					}
					if count2 == 22 {
						R22 = R31 - immediate
					}
					if count2 == 23 {
						R23 = R31 - immediate
					}
					if count2 == 24 {
						R24 = R31 - immediate
					}
					if count2 == 25 {
						R25 = R31 - immediate
					}
					if count2 == 26 {
						R26 = R31 - immediate
					}
					if count2 == 27 {
						R27 = R31 - immediate
					}
					if count2 == 28 {
						R28 = R31 - immediate
					}
					if count2 == 29 {
						R29 = R31 - immediate
					}
					if count2 == 30 {
						R30 = R31 - immediate
					}
					if count2 == 31 {
						R31 = R31 - immediate
					}
				}
			} else {
				//ADDI
				if count1 == 0 {
					if count2 == 0 {
						R0 = immediate + R0
					}
					if count2 == 1 {
						R1 = immediate + R0
					}
					if count2 == 2 {
						R2 = immediate + R0
					}
					if count2 == 3 {
						R3 = immediate + R0
					}
					if count2 == 4 {
						R4 = immediate + R0
					}
					if count2 == 5 {
						R5 = immediate + R0
					}
					if count2 == 6 {
						R6 = immediate + R0
					}
					if count2 == 7 {
						R7 = immediate + R0
					}
					if count2 == 8 {
						R8 = immediate + R0
					}
					if count2 == 9 {
						R9 = immediate + R0
					}
					if count2 == 10 {
						R10 = immediate + R0
					}
					if count2 == 11 {
						R11 = immediate + R0
					}
					if count2 == 12 {
						R12 = immediate + R0
					}
					if count2 == 13 {
						R13 = immediate + R0
					}
					if count2 == 14 {
						R14 = immediate + R0
					}
					if count2 == 15 {
						R15 = immediate + R0
					}
					if count2 == 16 {
						R16 = immediate + R0
					}
					if count2 == 17 {
						R17 = immediate + R0
					}
					if count2 == 18 {
						R18 = immediate + R0
					}
					if count2 == 19 {
						R19 = immediate + R0
					}
					if count2 == 20 {
						R20 = immediate + R0
					}
					if count2 == 21 {
						R21 = immediate + R0
					}
					if count2 == 22 {
						R22 = immediate + R0
					}
					if count2 == 23 {
						R23 = immediate + R0
					}
					if count2 == 24 {
						R24 = immediate + R0
					}
					if count2 == 25 {
						R25 = immediate + R0
					}
					if count2 == 26 {
						R26 = immediate + R0
					}
					if count2 == 27 {
						R27 = immediate + R0
					}
					if count2 == 28 {
						R28 = immediate + R0
					}
					if count2 == 29 {
						R29 = immediate + R0
					}
					if count2 == 30 {
						R30 = immediate + R0
					}
					if count2 == 31 {
						R31 = immediate + R0
					}
				}
				if count1 == 1 {
					if count2 == 0 {
						R0 = immediate + R1
					}
					if count2 == 1 {
						R1 = immediate + R1
					}
					if count2 == 2 {
						R2 = immediate + R1
					}
					if count2 == 3 {
						R3 = immediate + R1
					}
					if count2 == 4 {
						R4 = immediate + R1
					}
					if count2 == 5 {
						R5 = immediate + R1
					}
					if count2 == 6 {
						R6 = immediate + R1
					}
					if count2 == 7 {
						R7 = immediate + R1
					}
					if count2 == 8 {
						R8 = immediate + R1
					}
					if count2 == 9 {
						R9 = immediate + R1
					}
					if count2 == 10 {
						R10 = immediate + R1
					}
					if count2 == 11 {
						R11 = immediate + R1
					}
					if count2 == 12 {
						R12 = immediate + R1
					}
					if count2 == 13 {
						R13 = immediate + R1
					}
					if count2 == 14 {
						R14 = immediate + R1
					}
					if count2 == 15 {
						R15 = immediate + R1
					}
					if count2 == 16 {
						R16 = immediate + R1
					}
					if count2 == 17 {
						R17 = immediate + R1
					}
					if count2 == 18 {
						R18 = immediate + R1
					}
					if count2 == 19 {
						R19 = immediate + R1
					}
					if count2 == 20 {
						R20 = immediate + R1
					}
					if count2 == 21 {
						R21 = immediate + R1
					}
					if count2 == 22 {
						R22 = immediate + R1
					}
					if count2 == 23 {
						R23 = immediate + R1
					}
					if count2 == 24 {
						R24 = immediate + R1
					}
					if count2 == 25 {
						R25 = immediate + R1
					}
					if count2 == 26 {
						R26 = immediate + R1
					}
					if count2 == 27 {
						R27 = immediate + R1
					}
					if count2 == 28 {
						R28 = immediate + R1
					}
					if count2 == 29 {
						R29 = immediate + R1
					}
					if count2 == 30 {
						R30 = immediate + R1
					}
					if count2 == 31 {
						R31 = immediate + R1
					}
				}
				if count1 == 2 {
					if count2 == 0 {
						R0 = immediate + R2
					}
					if count2 == 1 {
						R1 = immediate + R2
					}
					if count2 == 2 {
						R2 = immediate + R2
					}
					if count2 == 3 {
						R3 = immediate + R2
					}
					if count2 == 4 {
						R4 = immediate + R2
					}
					if count2 == 5 {
						R5 = immediate + R2
					}
					if count2 == 6 {
						R6 = immediate + R2
					}
					if count2 == 7 {
						R7 = immediate + R2
					}
					if count2 == 8 {
						R8 = immediate + R2
					}
					if count2 == 9 {
						R9 = immediate + R2
					}
					if count2 == 10 {
						R10 = immediate + R2
					}
					if count2 == 11 {
						R11 = immediate + R2
					}
					if count2 == 12 {
						R12 = immediate + R2
					}
					if count2 == 13 {
						R13 = immediate + R2
					}
					if count2 == 14 {
						R14 = immediate + R2
					}
					if count2 == 15 {
						R15 = immediate + R2
					}
					if count2 == 16 {
						R16 = immediate + R2
					}
					if count2 == 17 {
						R17 = immediate + R2
					}
					if count2 == 18 {
						R18 = immediate + R2
					}
					if count2 == 19 {
						R19 = immediate + R2
					}
					if count2 == 20 {
						R20 = immediate + R2
					}
					if count2 == 21 {
						R21 = immediate + R2
					}
					if count2 == 22 {
						R22 = immediate + R2
					}
					if count2 == 23 {
						R23 = immediate + R2
					}
					if count2 == 24 {
						R24 = immediate + R2
					}
					if count2 == 25 {
						R25 = immediate + R2
					}
					if count2 == 26 {
						R26 = immediate + R2
					}
					if count2 == 27 {
						R27 = immediate + R2
					}
					if count2 == 28 {
						R28 = immediate + R2
					}
					if count2 == 29 {
						R29 = immediate + R2
					}
					if count2 == 30 {
						R30 = immediate + R2
					}
					if count2 == 31 {
						R31 = immediate + R2
					}
				}
				if count1 == 3 {
					if count2 == 0 {
						R0 = immediate + R3
					}
					if count2 == 1 {
						R1 = immediate + R3
					}
					if count2 == 2 {
						R2 = immediate + R3
					}
					if count2 == 3 {
						R3 = immediate + R3
					}
					if count2 == 4 {
						R4 = immediate + R3
					}
					if count2 == 5 {
						R5 = immediate + R3
					}
					if count2 == 6 {
						R6 = immediate + R3
					}
					if count2 == 7 {
						R7 = immediate + R3
					}
					if count2 == 8 {
						R8 = immediate + R3
					}
					if count2 == 9 {
						R9 = immediate + R3
					}
					if count2 == 10 {
						R10 = immediate + R3
					}
					if count2 == 11 {
						R11 = immediate + R3
					}
					if count2 == 12 {
						R12 = immediate + R3
					}
					if count2 == 13 {
						R13 = immediate + R3
					}
					if count2 == 14 {
						R14 = immediate + R3
					}
					if count2 == 15 {
						R15 = immediate + R3
					}
					if count2 == 16 {
						R16 = immediate + R3
					}
					if count2 == 17 {
						R17 = immediate + R3
					}
					if count2 == 18 {
						R18 = immediate + R3
					}
					if count2 == 19 {
						R19 = immediate + R3
					}
					if count2 == 20 {
						R20 = immediate + R3
					}
					if count2 == 21 {
						R21 = immediate + R3
					}
					if count2 == 22 {
						R22 = immediate + R3
					}
					if count2 == 23 {
						R23 = immediate + R3
					}
					if count2 == 24 {
						R24 = immediate + R3
					}
					if count2 == 25 {
						R25 = immediate + R3
					}
					if count2 == 26 {
						R26 = immediate + R3
					}
					if count2 == 27 {
						R27 = immediate + R3
					}
					if count2 == 28 {
						R28 = immediate + R3
					}
					if count2 == 29 {
						R29 = immediate + R3
					}
					if count2 == 30 {
						R30 = immediate + R3
					}
					if count2 == 31 {
						R31 = immediate + R3
					}
				}
				if count1 == 4 {
					if count2 == 0 {
						R0 = immediate + R4
					}
					if count2 == 1 {
						R1 = immediate + R4
					}
					if count2 == 2 {
						R2 = immediate + R4
					}
					if count2 == 3 {
						R3 = immediate + R4
					}
					if count2 == 4 {
						R4 = immediate + R4
					}
					if count2 == 5 {
						R5 = immediate + R4
					}
					if count2 == 6 {
						R6 = immediate + R4
					}
					if count2 == 7 {
						R7 = immediate + R4
					}
					if count2 == 8 {
						R8 = immediate + R4
					}
					if count2 == 9 {
						R9 = immediate + R4
					}
					if count2 == 10 {
						R10 = immediate + R4
					}
					if count2 == 11 {
						R11 = immediate + R4
					}
					if count2 == 12 {
						R12 = immediate + R4
					}
					if count2 == 13 {
						R13 = immediate + R4
					}
					if count2 == 14 {
						R14 = immediate + R4
					}
					if count2 == 15 {
						R15 = immediate + R4
					}
					if count2 == 16 {
						R16 = immediate + R4
					}
					if count2 == 17 {
						R17 = immediate + R4
					}
					if count2 == 18 {
						R18 = immediate + R4
					}
					if count2 == 19 {
						R19 = immediate + R4
					}
					if count2 == 20 {
						R20 = immediate + R4
					}
					if count2 == 21 {
						R21 = immediate + R4
					}
					if count2 == 22 {
						R22 = immediate + R4
					}
					if count2 == 23 {
						R23 = immediate + R4
					}
					if count2 == 24 {
						R24 = immediate + R4
					}
					if count2 == 25 {
						R25 = immediate + R4
					}
					if count2 == 26 {
						R26 = immediate + R4
					}
					if count2 == 27 {
						R27 = immediate + R4
					}
					if count2 == 28 {
						R28 = immediate + R4
					}
					if count2 == 29 {
						R29 = immediate + R4
					}
					if count2 == 30 {
						R30 = immediate + R4
					}
					if count2 == 31 {
						R31 = immediate + R4
					}
				}
				if count1 == 5 {
					if count2 == 0 {
						R0 = immediate + R5
					}
					if count2 == 1 {
						R1 = immediate + R5
					}
					if count2 == 2 {
						R2 = immediate + R5
					}
					if count2 == 3 {
						R3 = immediate + R5
					}
					if count2 == 4 {
						R4 = immediate + R5
					}
					if count2 == 5 {
						R5 = immediate + R5
					}
					if count2 == 6 {
						R6 = immediate + R5
					}
					if count2 == 7 {
						R7 = immediate + R5
					}
					if count2 == 8 {
						R8 = immediate + R5
					}
					if count2 == 9 {
						R9 = immediate + R5
					}
					if count2 == 10 {
						R10 = immediate + R5
					}
					if count2 == 11 {
						R11 = immediate + R5
					}
					if count2 == 12 {
						R12 = immediate + R5
					}
					if count2 == 13 {
						R13 = immediate + R5
					}
					if count2 == 14 {
						R14 = immediate + R5
					}
					if count2 == 15 {
						R15 = immediate + R5
					}
					if count2 == 16 {
						R16 = immediate + R5
					}
					if count2 == 17 {
						R17 = immediate + R5
					}
					if count2 == 18 {
						R18 = immediate + R5
					}
					if count2 == 19 {
						R19 = immediate + R5
					}
					if count2 == 20 {
						R20 = immediate + R5
					}
					if count2 == 21 {
						R21 = immediate + R5
					}
					if count2 == 22 {
						R22 = immediate + R5
					}
					if count2 == 23 {
						R23 = immediate + R5
					}
					if count2 == 24 {
						R24 = immediate + R5
					}
					if count2 == 25 {
						R25 = immediate + R5
					}
					if count2 == 26 {
						R26 = immediate + R5
					}
					if count2 == 27 {
						R27 = immediate + R5
					}
					if count2 == 28 {
						R28 = immediate + R5
					}
					if count2 == 29 {
						R29 = immediate + R5
					}
					if count2 == 30 {
						R30 = immediate + R5
					}
					if count2 == 31 {
						R31 = immediate + R5
					}
				}
				if count1 == 6 {
					if count2 == 0 {
						R0 = immediate + R6
					}
					if count2 == 1 {
						R1 = immediate + R6
					}
					if count2 == 2 {
						R2 = immediate + R6
					}
					if count2 == 3 {
						R3 = immediate + R6
					}
					if count2 == 4 {
						R4 = immediate + R6
					}
					if count2 == 5 {
						R5 = immediate + R6
					}
					if count2 == 6 {
						R6 = immediate + R6
					}
					if count2 == 7 {
						R7 = immediate + R6
					}
					if count2 == 8 {
						R8 = immediate + R6
					}
					if count2 == 9 {
						R9 = immediate + R6
					}
					if count2 == 10 {
						R10 = immediate + R6
					}
					if count2 == 11 {
						R11 = immediate + R6
					}
					if count2 == 12 {
						R12 = immediate + R6
					}
					if count2 == 13 {
						R13 = immediate + R6
					}
					if count2 == 14 {
						R14 = immediate + R6
					}
					if count2 == 15 {
						R15 = immediate + R6
					}
					if count2 == 16 {
						R16 = immediate + R6
					}
					if count2 == 17 {
						R17 = immediate + R6
					}
					if count2 == 18 {
						R18 = immediate + R6
					}
					if count2 == 19 {
						R19 = immediate + R6
					}
					if count2 == 20 {
						R20 = immediate + R6
					}
					if count2 == 21 {
						R21 = immediate + R6
					}
					if count2 == 22 {
						R22 = immediate + R6
					}
					if count2 == 23 {
						R23 = immediate + R6
					}
					if count2 == 24 {
						R24 = immediate + R6
					}
					if count2 == 25 {
						R25 = immediate + R6
					}
					if count2 == 26 {
						R26 = immediate + R6
					}
					if count2 == 27 {
						R27 = immediate + R6
					}
					if count2 == 28 {
						R28 = immediate + R6
					}
					if count2 == 29 {
						R29 = immediate + R6
					}
					if count2 == 30 {
						R30 = immediate + R6
					}
					if count2 == 31 {
						R31 = immediate + R6
					}
				}
				if count1 == 7 {
					if count2 == 0 {
						R0 = immediate + R7
					}
					if count2 == 1 {
						R1 = immediate + R7
					}
					if count2 == 2 {
						R2 = immediate + R7
					}
					if count2 == 3 {
						R3 = immediate + R7
					}
					if count2 == 4 {
						R4 = immediate + R7
					}
					if count2 == 5 {
						R5 = immediate + R7
					}
					if count2 == 6 {
						R6 = immediate + R7
					}
					if count2 == 7 {
						R7 = immediate + R7
					}
					if count2 == 8 {
						R8 = immediate + R7
					}
					if count2 == 9 {
						R9 = immediate + R7
					}
					if count2 == 10 {
						R10 = immediate + R7
					}
					if count2 == 11 {
						R11 = immediate + R7
					}
					if count2 == 12 {
						R12 = immediate + R7
					}
					if count2 == 13 {
						R13 = immediate + R7
					}
					if count2 == 14 {
						R14 = immediate + R7
					}
					if count2 == 15 {
						R15 = immediate + R7
					}
					if count2 == 16 {
						R16 = immediate + R7
					}
					if count2 == 17 {
						R17 = immediate + R7
					}
					if count2 == 18 {
						R18 = immediate + R7
					}
					if count2 == 19 {
						R19 = immediate + R7
					}
					if count2 == 20 {
						R20 = immediate + R7
					}
					if count2 == 21 {
						R21 = immediate + R7
					}
					if count2 == 22 {
						R22 = immediate + R7
					}
					if count2 == 23 {
						R23 = immediate + R7
					}
					if count2 == 24 {
						R24 = immediate + R7
					}
					if count2 == 25 {
						R25 = immediate + R7
					}
					if count2 == 26 {
						R26 = immediate + R7
					}
					if count2 == 27 {
						R27 = immediate + R7
					}
					if count2 == 28 {
						R28 = immediate + R7
					}
					if count2 == 29 {
						R29 = immediate + R7
					}
					if count2 == 30 {
						R30 = immediate + R7
					}
					if count2 == 31 {
						R31 = immediate + R7
					}
				}
				if count1 == 8 {
					if count2 == 0 {
						R0 = immediate + R8
					}
					if count2 == 1 {
						R1 = immediate + R8
					}
					if count2 == 2 {
						R2 = immediate + R8
					}
					if count2 == 3 {
						R3 = immediate + R8
					}
					if count2 == 4 {
						R4 = immediate + R8
					}
					if count2 == 5 {
						R5 = immediate + R8
					}
					if count2 == 6 {
						R6 = immediate + R8
					}
					if count2 == 7 {
						R7 = immediate + R8
					}
					if count2 == 8 {
						R8 = immediate + R8
					}
					if count2 == 9 {
						R9 = immediate + R8
					}
					if count2 == 10 {
						R10 = immediate + R8
					}
					if count2 == 11 {
						R11 = immediate + R8
					}
					if count2 == 12 {
						R12 = immediate + R8
					}
					if count2 == 13 {
						R13 = immediate + R8
					}
					if count2 == 14 {
						R14 = immediate + R8
					}
					if count2 == 15 {
						R15 = immediate + R8
					}
					if count2 == 16 {
						R16 = immediate + R8
					}
					if count2 == 17 {
						R17 = immediate + R8
					}
					if count2 == 18 {
						R18 = immediate + R8
					}
					if count2 == 19 {
						R19 = immediate + R8
					}
					if count2 == 20 {
						R20 = immediate + R8
					}
					if count2 == 21 {
						R21 = immediate + R8
					}
					if count2 == 22 {
						R22 = immediate + R8
					}
					if count2 == 23 {
						R23 = immediate + R8
					}
					if count2 == 24 {
						R24 = immediate + R8
					}
					if count2 == 25 {
						R25 = immediate + R8
					}
					if count2 == 26 {
						R26 = immediate + R8
					}
					if count2 == 27 {
						R27 = immediate + R8
					}
					if count2 == 28 {
						R28 = immediate + R8
					}
					if count2 == 29 {
						R29 = immediate + R8
					}
					if count2 == 30 {
						R30 = immediate + R8
					}
					if count2 == 31 {
						R31 = immediate + R8
					}
				}
				if count1 == 9 {
					if count2 == 0 {
						R0 = immediate + R9
					}
					if count2 == 1 {
						R1 = immediate + R9
					}
					if count2 == 2 {
						R2 = immediate + R9
					}
					if count2 == 3 {
						R3 = immediate + R9
					}
					if count2 == 4 {
						R4 = immediate + R9
					}
					if count2 == 5 {
						R5 = immediate + R9
					}
					if count2 == 6 {
						R6 = immediate + R9
					}
					if count2 == 7 {
						R7 = immediate + R9
					}
					if count2 == 8 {
						R8 = immediate + R9
					}
					if count2 == 9 {
						R9 = immediate + R9
					}
					if count2 == 10 {
						R10 = immediate + R9
					}
					if count2 == 11 {
						R11 = immediate + R9
					}
					if count2 == 12 {
						R12 = immediate + R9
					}
					if count2 == 13 {
						R13 = immediate + R9
					}
					if count2 == 14 {
						R14 = immediate + R9
					}
					if count2 == 15 {
						R15 = immediate + R9
					}
					if count2 == 16 {
						R16 = immediate + R9
					}
					if count2 == 17 {
						R17 = immediate + R9
					}
					if count2 == 18 {
						R18 = immediate + R9
					}
					if count2 == 19 {
						R19 = immediate + R9
					}
					if count2 == 20 {
						R20 = immediate + R9
					}
					if count2 == 21 {
						R21 = immediate + R9
					}
					if count2 == 22 {
						R22 = immediate + R9
					}
					if count2 == 23 {
						R23 = immediate + R9
					}
					if count2 == 24 {
						R24 = immediate + R9
					}
					if count2 == 25 {
						R25 = immediate + R9
					}
					if count2 == 26 {
						R26 = immediate + R9
					}
					if count2 == 27 {
						R27 = immediate + R9
					}
					if count2 == 28 {
						R28 = immediate + R9
					}
					if count2 == 29 {
						R29 = immediate + R9
					}
					if count2 == 30 {
						R30 = immediate + R9
					}
					if count2 == 31 {
						R31 = immediate + R9
					}
				}
				if count1 == 10 {
					if count2 == 0 {
						R0 = immediate + R10
					}
					if count2 == 1 {
						R1 = immediate + R10
					}
					if count2 == 2 {
						R2 = immediate + R10
					}
					if count2 == 3 {
						R3 = immediate + R10
					}
					if count2 == 4 {
						R4 = immediate + R10
					}
					if count2 == 5 {
						R5 = immediate + R10
					}
					if count2 == 6 {
						R6 = immediate + R10
					}
					if count2 == 7 {
						R7 = immediate + R10
					}
					if count2 == 8 {
						R8 = immediate + R10
					}
					if count2 == 9 {
						R9 = immediate + R10
					}
					if count2 == 10 {
						R10 = immediate + R10
					}
					if count2 == 11 {
						R11 = immediate + R10
					}
					if count2 == 12 {
						R12 = immediate + R10
					}
					if count2 == 13 {
						R13 = immediate + R10
					}
					if count2 == 14 {
						R14 = immediate + R10
					}
					if count2 == 15 {
						R15 = immediate + R10
					}
					if count2 == 16 {
						R16 = immediate + R10
					}
					if count2 == 17 {
						R17 = immediate + R10
					}
					if count2 == 18 {
						R18 = immediate + R10
					}
					if count2 == 19 {
						R19 = immediate + R10
					}
					if count2 == 20 {
						R20 = immediate + R10
					}
					if count2 == 21 {
						R21 = immediate + R10
					}
					if count2 == 22 {
						R22 = immediate + R10
					}
					if count2 == 23 {
						R23 = immediate + R10
					}
					if count2 == 24 {
						R24 = immediate + R10
					}
					if count2 == 25 {
						R25 = immediate + R10
					}
					if count2 == 26 {
						R26 = immediate + R10
					}
					if count2 == 27 {
						R27 = immediate + R10
					}
					if count2 == 28 {
						R28 = immediate + R10
					}
					if count2 == 29 {
						R29 = immediate + R10
					}
					if count2 == 30 {
						R30 = immediate + R10
					}
					if count2 == 31 {
						R31 = immediate + R10
					}
				}
				if count1 == 11 {
					if count2 == 0 {
						R0 = immediate + R11
					}
					if count2 == 1 {
						R1 = immediate + R11
					}
					if count2 == 2 {
						R2 = immediate + R11
					}
					if count2 == 3 {
						R3 = immediate + R11
					}
					if count2 == 4 {
						R4 = immediate + R11
					}
					if count2 == 5 {
						R5 = immediate + R11
					}
					if count2 == 6 {
						R6 = immediate + R11
					}
					if count2 == 7 {
						R7 = immediate + R11
					}
					if count2 == 8 {
						R8 = immediate + R11
					}
					if count2 == 9 {
						R9 = immediate + R11
					}
					if count2 == 10 {
						R10 = immediate + R11
					}
					if count2 == 11 {
						R11 = immediate + R11
					}
					if count2 == 12 {
						R12 = immediate + R11
					}
					if count2 == 13 {
						R13 = immediate + R11
					}
					if count2 == 14 {
						R14 = immediate + R11
					}
					if count2 == 15 {
						R15 = immediate + R11
					}
					if count2 == 16 {
						R16 = immediate + R11
					}
					if count2 == 17 {
						R17 = immediate + R11
					}
					if count2 == 18 {
						R18 = immediate + R11
					}
					if count2 == 19 {
						R19 = immediate + R11
					}
					if count2 == 20 {
						R20 = immediate + R11
					}
					if count2 == 21 {
						R21 = immediate + R11
					}
					if count2 == 22 {
						R22 = immediate + R11
					}
					if count2 == 23 {
						R23 = immediate + R11
					}
					if count2 == 24 {
						R24 = immediate + R11
					}
					if count2 == 25 {
						R25 = immediate + R11
					}
					if count2 == 26 {
						R26 = immediate + R11
					}
					if count2 == 27 {
						R27 = immediate + R11
					}
					if count2 == 28 {
						R28 = immediate + R11
					}
					if count2 == 29 {
						R29 = immediate + R11
					}
					if count2 == 30 {
						R30 = immediate + R11
					}
					if count2 == 31 {
						R31 = immediate + R11
					}
				}
				if count1 == 12 {
					if count2 == 0 {
						R0 = immediate + R12
					}
					if count2 == 1 {
						R1 = immediate + R12
					}
					if count2 == 2 {
						R2 = immediate + R12
					}
					if count2 == 3 {
						R3 = immediate + R12
					}
					if count2 == 4 {
						R4 = immediate + R12
					}
					if count2 == 5 {
						R5 = immediate + R12
					}
					if count2 == 6 {
						R6 = immediate + R12
					}
					if count2 == 7 {
						R7 = immediate + R12
					}
					if count2 == 8 {
						R8 = immediate + R12
					}
					if count2 == 9 {
						R9 = immediate + R12
					}
					if count2 == 10 {
						R10 = immediate + R12
					}
					if count2 == 11 {
						R11 = immediate + R12
					}
					if count2 == 12 {
						R12 = immediate + R12
					}
					if count2 == 13 {
						R13 = immediate + R12
					}
					if count2 == 14 {
						R14 = immediate + R12
					}
					if count2 == 15 {
						R15 = immediate + R12
					}
					if count2 == 16 {
						R16 = immediate + R12
					}
					if count2 == 17 {
						R17 = immediate + R12
					}
					if count2 == 18 {
						R18 = immediate + R12
					}
					if count2 == 19 {
						R19 = immediate + R12
					}
					if count2 == 20 {
						R20 = immediate + R12
					}
					if count2 == 21 {
						R21 = immediate + R12
					}
					if count2 == 22 {
						R22 = immediate + R12
					}
					if count2 == 23 {
						R23 = immediate + R12
					}
					if count2 == 24 {
						R24 = immediate + R12
					}
					if count2 == 25 {
						R25 = immediate + R12
					}
					if count2 == 26 {
						R26 = immediate + R12
					}
					if count2 == 27 {
						R27 = immediate + R12
					}
					if count2 == 28 {
						R28 = immediate + R12
					}
					if count2 == 29 {
						R29 = immediate + R12
					}
					if count2 == 30 {
						R30 = immediate + R12
					}
					if count2 == 31 {
						R31 = immediate + R12
					}
				}
				if count1 == 13 {
					if count2 == 0 {
						R0 = immediate + R13
					}
					if count2 == 1 {
						R1 = immediate + R13
					}
					if count2 == 2 {
						R2 = immediate + R13
					}
					if count2 == 3 {
						R3 = immediate + R13
					}
					if count2 == 4 {
						R4 = immediate + R13
					}
					if count2 == 5 {
						R5 = immediate + R13
					}
					if count2 == 6 {
						R6 = immediate + R13
					}
					if count2 == 7 {
						R7 = immediate + R13
					}
					if count2 == 8 {
						R8 = immediate + R13
					}
					if count2 == 9 {
						R9 = immediate + R13
					}
					if count2 == 10 {
						R10 = immediate + R13
					}
					if count2 == 11 {
						R11 = immediate + R13
					}
					if count2 == 12 {
						R12 = immediate + R13
					}
					if count2 == 13 {
						R13 = immediate + R13
					}
					if count2 == 14 {
						R14 = immediate + R13
					}
					if count2 == 15 {
						R15 = immediate + R13
					}
					if count2 == 16 {
						R16 = immediate + R13
					}
					if count2 == 17 {
						R17 = immediate + R13
					}
					if count2 == 18 {
						R18 = immediate + R13
					}
					if count2 == 19 {
						R19 = immediate + R13
					}
					if count2 == 20 {
						R20 = immediate + R13
					}
					if count2 == 21 {
						R21 = immediate + R13
					}
					if count2 == 22 {
						R22 = immediate + R13
					}
					if count2 == 23 {
						R23 = immediate + R13
					}
					if count2 == 24 {
						R24 = immediate + R13
					}
					if count2 == 25 {
						R25 = immediate + R13
					}
					if count2 == 26 {
						R26 = immediate + R13
					}
					if count2 == 27 {
						R27 = immediate + R13
					}
					if count2 == 28 {
						R28 = immediate + R13
					}
					if count2 == 29 {
						R29 = immediate + R13
					}
					if count2 == 30 {
						R30 = immediate + R13
					}
					if count2 == 31 {
						R31 = immediate + R13
					}
				}
				if count1 == 14 {
					if count2 == 0 {
						R0 = immediate + R14
					}
					if count2 == 1 {
						R1 = immediate + R14
					}
					if count2 == 2 {
						R2 = immediate + R14
					}
					if count2 == 3 {
						R3 = immediate + R14
					}
					if count2 == 4 {
						R4 = immediate + R14
					}
					if count2 == 5 {
						R5 = immediate + R14
					}
					if count2 == 6 {
						R6 = immediate + R14
					}
					if count2 == 7 {
						R7 = immediate + R14
					}
					if count2 == 8 {
						R8 = immediate + R14
					}
					if count2 == 9 {
						R9 = immediate + R14
					}
					if count2 == 10 {
						R10 = immediate + R14
					}
					if count2 == 11 {
						R11 = immediate + R14
					}
					if count2 == 12 {
						R12 = immediate + R14
					}
					if count2 == 13 {
						R13 = immediate + R14
					}
					if count2 == 14 {
						R14 = immediate + R14
					}
					if count2 == 15 {
						R15 = immediate + R14
					}
					if count2 == 16 {
						R16 = immediate + R14
					}
					if count2 == 17 {
						R17 = immediate + R14
					}
					if count2 == 18 {
						R18 = immediate + R14
					}
					if count2 == 19 {
						R19 = immediate + R14
					}
					if count2 == 20 {
						R20 = immediate + R14
					}
					if count2 == 21 {
						R21 = immediate + R14
					}
					if count2 == 22 {
						R22 = immediate + R14
					}
					if count2 == 23 {
						R23 = immediate + R14
					}
					if count2 == 24 {
						R24 = immediate + R14
					}
					if count2 == 25 {
						R25 = immediate + R14
					}
					if count2 == 26 {
						R26 = immediate + R14
					}
					if count2 == 27 {
						R27 = immediate + R14
					}
					if count2 == 28 {
						R28 = immediate + R14
					}
					if count2 == 29 {
						R29 = immediate + R14
					}
					if count2 == 30 {
						R30 = immediate + R14
					}
					if count2 == 31 {
						R31 = immediate + R14
					}
				}
				if count1 == 15 {
					if count2 == 0 {
						R0 = immediate + R15
					}
					if count2 == 1 {
						R1 = immediate + R15
					}
					if count2 == 2 {
						R2 = immediate + R15
					}
					if count2 == 3 {
						R3 = immediate + R15
					}
					if count2 == 4 {
						R4 = immediate + R15
					}
					if count2 == 5 {
						R5 = immediate + R15
					}
					if count2 == 6 {
						R6 = immediate + R15
					}
					if count2 == 7 {
						R7 = immediate + R15
					}
					if count2 == 8 {
						R8 = immediate + R15
					}
					if count2 == 9 {
						R9 = immediate + R15
					}
					if count2 == 10 {
						R10 = immediate + R15
					}
					if count2 == 11 {
						R11 = immediate + R15
					}
					if count2 == 12 {
						R12 = immediate + R15
					}
					if count2 == 13 {
						R13 = immediate + R15
					}
					if count2 == 14 {
						R14 = immediate + R15
					}
					if count2 == 15 {
						R15 = immediate + R15
					}
					if count2 == 16 {
						R16 = immediate + R15
					}
					if count2 == 17 {
						R17 = immediate + R15
					}
					if count2 == 18 {
						R18 = immediate + R15
					}
					if count2 == 19 {
						R19 = immediate + R15
					}
					if count2 == 20 {
						R20 = immediate + R15
					}
					if count2 == 21 {
						R21 = immediate + R15
					}
					if count2 == 22 {
						R22 = immediate + R15
					}
					if count2 == 23 {
						R23 = immediate + R15
					}
					if count2 == 24 {
						R24 = immediate + R15
					}
					if count2 == 25 {
						R25 = immediate + R15
					}
					if count2 == 26 {
						R26 = immediate + R15
					}
					if count2 == 27 {
						R27 = immediate + R15
					}
					if count2 == 28 {
						R28 = immediate + R15
					}
					if count2 == 29 {
						R29 = immediate + R15
					}
					if count2 == 30 {
						R30 = immediate + R15
					}
					if count2 == 31 {
						R31 = immediate + R15
					}
				}
				if count1 == 16 {
					if count2 == 0 {
						R0 = immediate + R16
					}
					if count2 == 1 {
						R1 = immediate + R16
					}
					if count2 == 2 {
						R2 = immediate + R16
					}
					if count2 == 3 {
						R3 = immediate + R16
					}
					if count2 == 4 {
						R4 = immediate + R16
					}
					if count2 == 5 {
						R5 = immediate + R16
					}
					if count2 == 6 {
						R6 = immediate + R16
					}
					if count2 == 7 {
						R7 = immediate + R16
					}
					if count2 == 8 {
						R8 = immediate + R16
					}
					if count2 == 9 {
						R9 = immediate + R16
					}
					if count2 == 10 {
						R10 = immediate + R16
					}
					if count2 == 11 {
						R11 = immediate + R16
					}
					if count2 == 12 {
						R12 = immediate + R16
					}
					if count2 == 13 {
						R13 = immediate + R16
					}
					if count2 == 14 {
						R14 = immediate + R16
					}
					if count2 == 15 {
						R15 = immediate + R16
					}
					if count2 == 16 {
						R16 = immediate + R16
					}
					if count2 == 17 {
						R17 = immediate + R16
					}
					if count2 == 18 {
						R18 = immediate + R16
					}
					if count2 == 19 {
						R19 = immediate + R16
					}
					if count2 == 20 {
						R20 = immediate + R16
					}
					if count2 == 21 {
						R21 = immediate + R16
					}
					if count2 == 22 {
						R22 = immediate + R16
					}
					if count2 == 23 {
						R23 = immediate + R16
					}
					if count2 == 24 {
						R24 = immediate + R16
					}
					if count2 == 25 {
						R25 = immediate + R16
					}
					if count2 == 26 {
						R26 = immediate + R16
					}
					if count2 == 27 {
						R27 = immediate + R16
					}
					if count2 == 28 {
						R28 = immediate + R16
					}
					if count2 == 29 {
						R29 = immediate + R16
					}
					if count2 == 30 {
						R30 = immediate + R16
					}
					if count2 == 31 {
						R31 = immediate + R16
					}
				}
				if count1 == 17 {
					if count2 == 0 {
						R0 = immediate + R17
					}
					if count2 == 1 {
						R1 = immediate + R17
					}
					if count2 == 2 {
						R2 = immediate + R17
					}
					if count2 == 3 {
						R3 = immediate + R17
					}
					if count2 == 4 {
						R4 = immediate + R17
					}
					if count2 == 5 {
						R5 = immediate + R17
					}
					if count2 == 6 {
						R6 = immediate + R17
					}
					if count2 == 7 {
						R7 = immediate + R17
					}
					if count2 == 8 {
						R8 = immediate + R17
					}
					if count2 == 9 {
						R9 = immediate + R17
					}
					if count2 == 10 {
						R10 = immediate + R17
					}
					if count2 == 11 {
						R11 = immediate + R17
					}
					if count2 == 12 {
						R12 = immediate + R17
					}
					if count2 == 13 {
						R13 = immediate + R17
					}
					if count2 == 14 {
						R14 = immediate + R17
					}
					if count2 == 15 {
						R15 = immediate + R17
					}
					if count2 == 16 {
						R16 = immediate + R17
					}
					if count2 == 17 {
						R17 = immediate + R17
					}
					if count2 == 18 {
						R18 = immediate + R17
					}
					if count2 == 19 {
						R19 = immediate + R17
					}
					if count2 == 20 {
						R20 = immediate + R17
					}
					if count2 == 21 {
						R21 = immediate + R17
					}
					if count2 == 22 {
						R22 = immediate + R17
					}
					if count2 == 23 {
						R23 = immediate + R17
					}
					if count2 == 24 {
						R24 = immediate + R17
					}
					if count2 == 25 {
						R25 = immediate + R17
					}
					if count2 == 26 {
						R26 = immediate + R17
					}
					if count2 == 27 {
						R27 = immediate + R17
					}
					if count2 == 28 {
						R28 = immediate + R17
					}
					if count2 == 29 {
						R29 = immediate + R17
					}
					if count2 == 30 {
						R30 = immediate + R17
					}
					if count2 == 31 {
						R31 = immediate + R17
					}
				}
				if count1 == 18 {
					if count2 == 0 {
						R0 = immediate + R18
					}
					if count2 == 1 {
						R1 = immediate + R18
					}
					if count2 == 2 {
						R2 = immediate + R18
					}
					if count2 == 3 {
						R3 = immediate + R18
					}
					if count2 == 4 {
						R4 = immediate + R18
					}
					if count2 == 5 {
						R5 = immediate + R18
					}
					if count2 == 6 {
						R6 = immediate + R18
					}
					if count2 == 7 {
						R7 = immediate + R18
					}
					if count2 == 8 {
						R8 = immediate + R18
					}
					if count2 == 9 {
						R9 = immediate + R18
					}
					if count2 == 10 {
						R10 = immediate + R18
					}
					if count2 == 11 {
						R11 = immediate + R18
					}
					if count2 == 12 {
						R12 = immediate + R18
					}
					if count2 == 13 {
						R13 = immediate + R18
					}
					if count2 == 14 {
						R14 = immediate + R18
					}
					if count2 == 15 {
						R15 = immediate + R18
					}
					if count2 == 16 {
						R16 = immediate + R18
					}
					if count2 == 17 {
						R17 = immediate + R18
					}
					if count2 == 18 {
						R18 = immediate + R18
					}
					if count2 == 19 {
						R19 = immediate + R18
					}
					if count2 == 20 {
						R20 = immediate + R18
					}
					if count2 == 21 {
						R21 = immediate + R18
					}
					if count2 == 22 {
						R22 = immediate + R18
					}
					if count2 == 23 {
						R23 = immediate + R18
					}
					if count2 == 24 {
						R24 = immediate + R18
					}
					if count2 == 25 {
						R25 = immediate + R18
					}
					if count2 == 26 {
						R26 = immediate + R18
					}
					if count2 == 27 {
						R27 = immediate + R18
					}
					if count2 == 28 {
						R28 = immediate + R18
					}
					if count2 == 29 {
						R29 = immediate + R18
					}
					if count2 == 30 {
						R30 = immediate + R18
					}
					if count2 == 31 {
						R31 = immediate + R18
					}
				}
				if count1 == 19 {
					if count2 == 0 {
						R0 = immediate + R19
					}
					if count2 == 1 {
						R1 = immediate + R19
					}
					if count2 == 2 {
						R2 = immediate + R19
					}
					if count2 == 3 {
						R3 = immediate + R19
					}
					if count2 == 4 {
						R4 = immediate + R19
					}
					if count2 == 5 {
						R5 = immediate + R19
					}
					if count2 == 6 {
						R6 = immediate + R19
					}
					if count2 == 7 {
						R7 = immediate + R19
					}
					if count2 == 8 {
						R8 = immediate + R19
					}
					if count2 == 9 {
						R9 = immediate + R19
					}
					if count2 == 10 {
						R10 = immediate + R19
					}
					if count2 == 11 {
						R11 = immediate + R19
					}
					if count2 == 12 {
						R12 = immediate + R19
					}
					if count2 == 13 {
						R13 = immediate + R19
					}
					if count2 == 14 {
						R14 = immediate + R19
					}
					if count2 == 15 {
						R15 = immediate + R19
					}
					if count2 == 16 {
						R16 = immediate + R19
					}
					if count2 == 17 {
						R17 = immediate + R19
					}
					if count2 == 18 {
						R18 = immediate + R19
					}
					if count2 == 19 {
						R19 = immediate + R19
					}
					if count2 == 20 {
						R20 = immediate + R19
					}
					if count2 == 21 {
						R21 = immediate + R19
					}
					if count2 == 22 {
						R22 = immediate + R19
					}
					if count2 == 23 {
						R23 = immediate + R19
					}
					if count2 == 24 {
						R24 = immediate + R19
					}
					if count2 == 25 {
						R25 = immediate + R19
					}
					if count2 == 26 {
						R26 = immediate + R19
					}
					if count2 == 27 {
						R27 = immediate + R19
					}
					if count2 == 28 {
						R28 = immediate + R19
					}
					if count2 == 29 {
						R29 = immediate + R19
					}
					if count2 == 30 {
						R30 = immediate + R19
					}
					if count2 == 31 {
						R31 = immediate + R19
					}
				}
				if count1 == 20 {
					if count2 == 0 {
						R0 = immediate + R20
					}
					if count2 == 1 {
						R1 = immediate + R20
					}
					if count2 == 2 {
						R2 = immediate + R20
					}
					if count2 == 3 {
						R3 = immediate + R20
					}
					if count2 == 4 {
						R4 = immediate + R20
					}
					if count2 == 5 {
						R5 = immediate + R20
					}
					if count2 == 6 {
						R6 = immediate + R20
					}
					if count2 == 7 {
						R7 = immediate + R20
					}
					if count2 == 8 {
						R8 = immediate + R20
					}
					if count2 == 9 {
						R9 = immediate + R20
					}
					if count2 == 10 {
						R10 = immediate + R20
					}
					if count2 == 11 {
						R11 = immediate + R20
					}
					if count2 == 12 {
						R12 = immediate + R20
					}
					if count2 == 13 {
						R13 = immediate + R20
					}
					if count2 == 14 {
						R14 = immediate + R20
					}
					if count2 == 15 {
						R15 = immediate + R20
					}
					if count2 == 16 {
						R16 = immediate + R20
					}
					if count2 == 17 {
						R17 = immediate + R20
					}
					if count2 == 18 {
						R18 = immediate + R20
					}
					if count2 == 19 {
						R19 = immediate + R20
					}
					if count2 == 20 {
						R20 = immediate + R20
					}
					if count2 == 21 {
						R21 = immediate + R20
					}
					if count2 == 22 {
						R22 = immediate + R20
					}
					if count2 == 23 {
						R23 = immediate + R20
					}
					if count2 == 24 {
						R24 = immediate + R20
					}
					if count2 == 25 {
						R25 = immediate + R20
					}
					if count2 == 26 {
						R26 = immediate + R20
					}
					if count2 == 27 {
						R27 = immediate + R20
					}
					if count2 == 28 {
						R28 = immediate + R20
					}
					if count2 == 29 {
						R29 = immediate + R20
					}
					if count2 == 30 {
						R30 = immediate + R20
					}
					if count2 == 31 {
						R31 = immediate + R20
					}
				}
				if count1 == 21 {
					if count2 == 0 {
						R0 = immediate + R21
					}
					if count2 == 1 {
						R1 = immediate + R21
					}
					if count2 == 2 {
						R2 = immediate + R21
					}
					if count2 == 3 {
						R3 = immediate + R21
					}
					if count2 == 4 {
						R4 = immediate + R21
					}
					if count2 == 5 {
						R5 = immediate + R21
					}
					if count2 == 6 {
						R6 = immediate + R21
					}
					if count2 == 7 {
						R7 = immediate + R21
					}
					if count2 == 8 {
						R8 = immediate + R21
					}
					if count2 == 9 {
						R9 = immediate + R21
					}
					if count2 == 10 {
						R10 = immediate + R21
					}
					if count2 == 11 {
						R11 = immediate + R21
					}
					if count2 == 12 {
						R12 = immediate + R21
					}
					if count2 == 13 {
						R13 = immediate + R21
					}
					if count2 == 14 {
						R14 = immediate + R21
					}
					if count2 == 15 {
						R15 = immediate + R21
					}
					if count2 == 16 {
						R16 = immediate + R21
					}
					if count2 == 17 {
						R17 = immediate + R21
					}
					if count2 == 18 {
						R18 = immediate + R21
					}
					if count2 == 19 {
						R19 = immediate + R21
					}
					if count2 == 20 {
						R20 = immediate + R21
					}
					if count2 == 21 {
						R21 = immediate + R21
					}
					if count2 == 22 {
						R22 = immediate + R21
					}
					if count2 == 23 {
						R23 = immediate + R21
					}
					if count2 == 24 {
						R24 = immediate + R21
					}
					if count2 == 25 {
						R25 = immediate + R21
					}
					if count2 == 26 {
						R26 = immediate + R21
					}
					if count2 == 27 {
						R27 = immediate + R21
					}
					if count2 == 28 {
						R28 = immediate + R21
					}
					if count2 == 29 {
						R29 = immediate + R21
					}
					if count2 == 30 {
						R30 = immediate + R21
					}
					if count2 == 31 {
						R31 = immediate + R21
					}
				}
				if count1 == 22 {
					if count2 == 0 {
						R0 = immediate + R22
					}
					if count2 == 1 {
						R1 = immediate + R22
					}
					if count2 == 2 {
						R2 = immediate + R22
					}
					if count2 == 3 {
						R3 = immediate + R22
					}
					if count2 == 4 {
						R4 = immediate + R22
					}
					if count2 == 5 {
						R5 = immediate + R22
					}
					if count2 == 6 {
						R6 = immediate + R22
					}
					if count2 == 7 {
						R7 = immediate + R22
					}
					if count2 == 8 {
						R8 = immediate + R22
					}
					if count2 == 9 {
						R9 = immediate + R22
					}
					if count2 == 10 {
						R10 = immediate + R22
					}
					if count2 == 11 {
						R11 = immediate + R22
					}
					if count2 == 12 {
						R12 = immediate + R22
					}
					if count2 == 13 {
						R13 = immediate + R22
					}
					if count2 == 14 {
						R14 = immediate + R22
					}
					if count2 == 15 {
						R15 = immediate + R22
					}
					if count2 == 16 {
						R16 = immediate + R22
					}
					if count2 == 17 {
						R17 = immediate + R22
					}
					if count2 == 18 {
						R18 = immediate + R22
					}
					if count2 == 19 {
						R19 = immediate + R22
					}
					if count2 == 20 {
						R20 = immediate + R22
					}
					if count2 == 21 {
						R21 = immediate + R22
					}
					if count2 == 22 {
						R22 = immediate + R22
					}
					if count2 == 23 {
						R23 = immediate + R22
					}
					if count2 == 24 {
						R24 = immediate + R22
					}
					if count2 == 25 {
						R25 = immediate + R22
					}
					if count2 == 26 {
						R26 = immediate + R22
					}
					if count2 == 27 {
						R27 = immediate + R22
					}
					if count2 == 28 {
						R28 = immediate + R22
					}
					if count2 == 29 {
						R29 = immediate + R22
					}
					if count2 == 30 {
						R30 = immediate + R22
					}
					if count2 == 31 {
						R31 = immediate + R22
					}
				}
				if count1 == 23 {
					if count2 == 0 {
						R0 = immediate + R23
					}
					if count2 == 1 {
						R1 = immediate + R23
					}
					if count2 == 2 {
						R2 = immediate + R23
					}
					if count2 == 3 {
						R3 = immediate + R23
					}
					if count2 == 4 {
						R4 = immediate + R23
					}
					if count2 == 5 {
						R5 = immediate + R23
					}
					if count2 == 6 {
						R6 = immediate + R23
					}
					if count2 == 7 {
						R7 = immediate + R23
					}
					if count2 == 8 {
						R8 = immediate + R23
					}
					if count2 == 9 {
						R9 = immediate + R23
					}
					if count2 == 10 {
						R10 = immediate + R23
					}
					if count2 == 11 {
						R11 = immediate + R23
					}
					if count2 == 12 {
						R12 = immediate + R23
					}
					if count2 == 13 {
						R13 = immediate + R23
					}
					if count2 == 14 {
						R14 = immediate + R23
					}
					if count2 == 15 {
						R15 = immediate + R23
					}
					if count2 == 16 {
						R16 = immediate + R23
					}
					if count2 == 17 {
						R17 = immediate + R23
					}
					if count2 == 18 {
						R18 = immediate + R23
					}
					if count2 == 19 {
						R19 = immediate + R23
					}
					if count2 == 20 {
						R20 = immediate + R23
					}
					if count2 == 21 {
						R21 = immediate + R23
					}
					if count2 == 22 {
						R22 = immediate + R23
					}
					if count2 == 23 {
						R23 = immediate + R23
					}
					if count2 == 24 {
						R24 = immediate + R23
					}
					if count2 == 25 {
						R25 = immediate + R23
					}
					if count2 == 26 {
						R26 = immediate + R23
					}
					if count2 == 27 {
						R27 = immediate + R23
					}
					if count2 == 28 {
						R28 = immediate + R23
					}
					if count2 == 29 {
						R29 = immediate + R23
					}
					if count2 == 30 {
						R30 = immediate + R23
					}
					if count2 == 31 {
						R31 = immediate + R23
					}
				}
				if count1 == 24 {
					if count2 == 0 {
						R0 = immediate + R24
					}
					if count2 == 1 {
						R1 = immediate + R24
					}
					if count2 == 2 {
						R2 = immediate + R24
					}
					if count2 == 3 {
						R3 = immediate + R24
					}
					if count2 == 4 {
						R4 = immediate + R24
					}
					if count2 == 5 {
						R5 = immediate + R24
					}
					if count2 == 6 {
						R6 = immediate + R24
					}
					if count2 == 7 {
						R7 = immediate + R24
					}
					if count2 == 8 {
						R8 = immediate + R24
					}
					if count2 == 9 {
						R9 = immediate + R24
					}
					if count2 == 10 {
						R10 = immediate + R24
					}
					if count2 == 11 {
						R11 = immediate + R24
					}
					if count2 == 12 {
						R12 = immediate + R24
					}
					if count2 == 13 {
						R13 = immediate + R24
					}
					if count2 == 14 {
						R14 = immediate + R24
					}
					if count2 == 15 {
						R15 = immediate + R24
					}
					if count2 == 16 {
						R16 = immediate + R24
					}
					if count2 == 17 {
						R17 = immediate + R24
					}
					if count2 == 18 {
						R18 = immediate + R24
					}
					if count2 == 19 {
						R19 = immediate + R24
					}
					if count2 == 20 {
						R20 = immediate + R24
					}
					if count2 == 21 {
						R21 = immediate + R24
					}
					if count2 == 22 {
						R22 = immediate + R24
					}
					if count2 == 23 {
						R23 = immediate + R24
					}
					if count2 == 24 {
						R24 = immediate + R24
					}
					if count2 == 25 {
						R25 = immediate + R24
					}
					if count2 == 26 {
						R26 = immediate + R24
					}
					if count2 == 27 {
						R27 = immediate + R24
					}
					if count2 == 28 {
						R28 = immediate + R24
					}
					if count2 == 29 {
						R29 = immediate + R24
					}
					if count2 == 30 {
						R30 = immediate + R24
					}
					if count2 == 31 {
						R31 = immediate + R24
					}
				}
				if count1 == 25 {
					if count2 == 0 {
						R0 = immediate + R25
					}
					if count2 == 1 {
						R1 = immediate + R25
					}
					if count2 == 2 {
						R2 = immediate + R25
					}
					if count2 == 3 {
						R3 = immediate + R25
					}
					if count2 == 4 {
						R4 = immediate + R25
					}
					if count2 == 5 {
						R5 = immediate + R25
					}
					if count2 == 6 {
						R6 = immediate + R25
					}
					if count2 == 7 {
						R7 = immediate + R25
					}
					if count2 == 8 {
						R8 = immediate + R25
					}
					if count2 == 9 {
						R9 = immediate + R25
					}
					if count2 == 10 {
						R10 = immediate + R25
					}
					if count2 == 11 {
						R11 = immediate + R25
					}
					if count2 == 12 {
						R12 = immediate + R25
					}
					if count2 == 13 {
						R13 = immediate + R25
					}
					if count2 == 14 {
						R14 = immediate + R25
					}
					if count2 == 15 {
						R15 = immediate + R25
					}
					if count2 == 16 {
						R16 = immediate + R25
					}
					if count2 == 17 {
						R17 = immediate + R25
					}
					if count2 == 18 {
						R18 = immediate + R25
					}
					if count2 == 19 {
						R19 = immediate + R25
					}
					if count2 == 20 {
						R20 = immediate + R25
					}
					if count2 == 21 {
						R21 = immediate + R25
					}
					if count2 == 22 {
						R22 = immediate + R25
					}
					if count2 == 23 {
						R23 = immediate + R25
					}
					if count2 == 24 {
						R24 = immediate + R25
					}
					if count2 == 25 {
						R25 = immediate + R25
					}
					if count2 == 26 {
						R26 = immediate + R25
					}
					if count2 == 27 {
						R27 = immediate + R25
					}
					if count2 == 28 {
						R28 = immediate + R25
					}
					if count2 == 29 {
						R29 = immediate + R25
					}
					if count2 == 30 {
						R30 = immediate + R25
					}
					if count2 == 31 {
						R31 = immediate + R25
					}
				}
				if count1 == 26 {
					if count2 == 0 {
						R0 = immediate + R26
					}
					if count2 == 1 {
						R1 = immediate + R26
					}
					if count2 == 2 {
						R2 = immediate + R26
					}
					if count2 == 3 {
						R3 = immediate + R26
					}
					if count2 == 4 {
						R4 = immediate + R26
					}
					if count2 == 5 {
						R5 = immediate + R26
					}
					if count2 == 6 {
						R6 = immediate + R26
					}
					if count2 == 7 {
						R7 = immediate + R26
					}
					if count2 == 8 {
						R8 = immediate + R26
					}
					if count2 == 9 {
						R9 = immediate + R26
					}
					if count2 == 10 {
						R10 = immediate + R26
					}
					if count2 == 11 {
						R11 = immediate + R26
					}
					if count2 == 12 {
						R12 = immediate + R26
					}
					if count2 == 13 {
						R13 = immediate + R26
					}
					if count2 == 14 {
						R14 = immediate + R26
					}
					if count2 == 15 {
						R15 = immediate + R26
					}
					if count2 == 16 {
						R16 = immediate + R26
					}
					if count2 == 17 {
						R17 = immediate + R26
					}
					if count2 == 18 {
						R18 = immediate + R26
					}
					if count2 == 19 {
						R19 = immediate + R26
					}
					if count2 == 20 {
						R20 = immediate + R26
					}
					if count2 == 21 {
						R21 = immediate + R26
					}
					if count2 == 22 {
						R22 = immediate + R26
					}
					if count2 == 23 {
						R23 = immediate + R26
					}
					if count2 == 24 {
						R24 = immediate + R26
					}
					if count2 == 25 {
						R25 = immediate + R26
					}
					if count2 == 26 {
						R26 = immediate + R26
					}
					if count2 == 27 {
						R27 = immediate + R26
					}
					if count2 == 28 {
						R28 = immediate + R26
					}
					if count2 == 29 {
						R29 = immediate + R26
					}
					if count2 == 30 {
						R30 = immediate + R26
					}
					if count2 == 31 {
						R31 = immediate + R26
					}
				}
				if count1 == 27 {
					if count2 == 0 {
						R0 = immediate + R27
					}
					if count2 == 1 {
						R1 = immediate + R27
					}
					if count2 == 2 {
						R2 = immediate + R27
					}
					if count2 == 3 {
						R3 = immediate + R27
					}
					if count2 == 4 {
						R4 = immediate + R27
					}
					if count2 == 5 {
						R5 = immediate + R27
					}
					if count2 == 6 {
						R6 = immediate + R27
					}
					if count2 == 7 {
						R7 = immediate + R27
					}
					if count2 == 8 {
						R8 = immediate + R27
					}
					if count2 == 9 {
						R9 = immediate + R27
					}
					if count2 == 10 {
						R10 = immediate + R27
					}
					if count2 == 11 {
						R11 = immediate + R27
					}
					if count2 == 12 {
						R12 = immediate + R27
					}
					if count2 == 13 {
						R13 = immediate + R27
					}
					if count2 == 14 {
						R14 = immediate + R27
					}
					if count2 == 15 {
						R15 = immediate + R27
					}
					if count2 == 16 {
						R16 = immediate + R27
					}
					if count2 == 17 {
						R17 = immediate + R27
					}
					if count2 == 18 {
						R18 = immediate + R27
					}
					if count2 == 19 {
						R19 = immediate + R27
					}
					if count2 == 20 {
						R20 = immediate + R27
					}
					if count2 == 21 {
						R21 = immediate + R27
					}
					if count2 == 22 {
						R22 = immediate + R27
					}
					if count2 == 23 {
						R23 = immediate + R27
					}
					if count2 == 24 {
						R24 = immediate + R27
					}
					if count2 == 25 {
						R25 = immediate + R27
					}
					if count2 == 26 {
						R26 = immediate + R27
					}
					if count2 == 27 {
						R27 = immediate + R27
					}
					if count2 == 28 {
						R28 = immediate + R27
					}
					if count2 == 29 {
						R29 = immediate + R27
					}
					if count2 == 30 {
						R30 = immediate + R27
					}
					if count2 == 31 {
						R31 = immediate + R27
					}
				}
				if count1 == 28 {
					if count2 == 0 {
						R0 = immediate + R28
					}
					if count2 == 1 {
						R1 = immediate + R28
					}
					if count2 == 2 {
						R2 = immediate + R28
					}
					if count2 == 3 {
						R3 = immediate + R28
					}
					if count2 == 4 {
						R4 = immediate + R28
					}
					if count2 == 5 {
						R5 = immediate + R28
					}
					if count2 == 6 {
						R6 = immediate + R28
					}
					if count2 == 7 {
						R7 = immediate + R28
					}
					if count2 == 8 {
						R8 = immediate + R28
					}
					if count2 == 9 {
						R9 = immediate + R28
					}
					if count2 == 10 {
						R10 = immediate + R28
					}
					if count2 == 11 {
						R11 = immediate + R28
					}
					if count2 == 12 {
						R12 = immediate + R28
					}
					if count2 == 13 {
						R13 = immediate + R28
					}
					if count2 == 14 {
						R14 = immediate + R28
					}
					if count2 == 15 {
						R15 = immediate + R28
					}
					if count2 == 16 {
						R16 = immediate + R28
					}
					if count2 == 17 {
						R17 = immediate + R28
					}
					if count2 == 18 {
						R18 = immediate + R28
					}
					if count2 == 19 {
						R19 = immediate + R28
					}
					if count2 == 20 {
						R20 = immediate + R28
					}
					if count2 == 21 {
						R21 = immediate + R28
					}
					if count2 == 22 {
						R22 = immediate + R28
					}
					if count2 == 23 {
						R23 = immediate + R28
					}
					if count2 == 24 {
						R24 = immediate + R28
					}
					if count2 == 25 {
						R25 = immediate + R28
					}
					if count2 == 26 {
						R26 = immediate + R28
					}
					if count2 == 27 {
						R27 = immediate + R28
					}
					if count2 == 28 {
						R28 = immediate + R28
					}
					if count2 == 29 {
						R29 = immediate + R28
					}
					if count2 == 30 {
						R30 = immediate + R28
					}
					if count2 == 31 {
						R31 = immediate + R28
					}
				}
				if count1 == 29 {
					if count2 == 0 {
						R0 = immediate + R29
					}
					if count2 == 1 {
						R1 = immediate + R29
					}
					if count2 == 2 {
						R2 = immediate + R29
					}
					if count2 == 3 {
						R3 = immediate + R29
					}
					if count2 == 4 {
						R4 = immediate + R29
					}
					if count2 == 5 {
						R5 = immediate + R29
					}
					if count2 == 6 {
						R6 = immediate + R29
					}
					if count2 == 7 {
						R7 = immediate + R29
					}
					if count2 == 8 {
						R8 = immediate + R29
					}
					if count2 == 9 {
						R9 = immediate + R29
					}
					if count2 == 10 {
						R10 = immediate + R29
					}
					if count2 == 11 {
						R11 = immediate + R29
					}
					if count2 == 12 {
						R12 = immediate + R29
					}
					if count2 == 13 {
						R13 = immediate + R29
					}
					if count2 == 14 {
						R14 = immediate + R29
					}
					if count2 == 15 {
						R15 = immediate + R29
					}
					if count2 == 16 {
						R16 = immediate + R29
					}
					if count2 == 17 {
						R17 = immediate + R29
					}
					if count2 == 18 {
						R18 = immediate + R29
					}
					if count2 == 19 {
						R19 = immediate + R29
					}
					if count2 == 20 {
						R20 = immediate + R29
					}
					if count2 == 21 {
						R21 = immediate + R29
					}
					if count2 == 22 {
						R22 = immediate + R29
					}
					if count2 == 23 {
						R23 = immediate + R29
					}
					if count2 == 24 {
						R24 = immediate + R29
					}
					if count2 == 25 {
						R25 = immediate + R29
					}
					if count2 == 26 {
						R26 = immediate + R29
					}
					if count2 == 27 {
						R27 = immediate + R29
					}
					if count2 == 28 {
						R28 = immediate + R29
					}
					if count2 == 29 {
						R29 = immediate + R29
					}
					if count2 == 30 {
						R30 = immediate + R29
					}
					if count2 == 31 {
						R31 = immediate + R29
					}
				}
				if count1 == 30 {
					if count2 == 0 {
						R0 = immediate + R30
					}
					if count2 == 1 {
						R1 = immediate + R30
					}
					if count2 == 2 {
						R2 = immediate + R30
					}
					if count2 == 3 {
						R3 = immediate + R30
					}
					if count2 == 4 {
						R4 = immediate + R30
					}
					if count2 == 5 {
						R5 = immediate + R30
					}
					if count2 == 6 {
						R6 = immediate + R30
					}
					if count2 == 7 {
						R7 = immediate + R30
					}
					if count2 == 8 {
						R8 = immediate + R30
					}
					if count2 == 9 {
						R9 = immediate + R30
					}
					if count2 == 10 {
						R10 = immediate + R30
					}
					if count2 == 11 {
						R11 = immediate + R30
					}
					if count2 == 12 {
						R12 = immediate + R30
					}
					if count2 == 13 {
						R13 = immediate + R30
					}
					if count2 == 14 {
						R14 = immediate + R30
					}
					if count2 == 15 {
						R15 = immediate + R30
					}
					if count2 == 16 {
						R16 = immediate + R30
					}
					if count2 == 17 {
						R17 = immediate + R30
					}
					if count2 == 18 {
						R18 = immediate + R30
					}
					if count2 == 19 {
						R19 = immediate + R30
					}
					if count2 == 20 {
						R20 = immediate + R30
					}
					if count2 == 21 {
						R21 = immediate + R30
					}
					if count2 == 22 {
						R22 = immediate + R30
					}
					if count2 == 23 {
						R23 = immediate + R30
					}
					if count2 == 24 {
						R24 = immediate + R30
					}
					if count2 == 25 {
						R25 = immediate + R30
					}
					if count2 == 26 {
						R26 = immediate + R30
					}
					if count2 == 27 {
						R27 = immediate + R30
					}
					if count2 == 28 {
						R28 = immediate + R30
					}
					if count2 == 29 {
						R29 = immediate + R30
					}
					if count2 == 30 {
						R30 = immediate + R30
					}
					if count2 == 31 {
						R31 = immediate + R30
					}
				}
				if count1 == 31 {
					if count2 == 0 {
						R0 = immediate + R31
					}
					if count2 == 1 {
						R1 = immediate + R31
					}
					if count2 == 2 {
						R2 = immediate + R31
					}
					if count2 == 3 {
						R3 = immediate + R31
					}
					if count2 == 4 {
						R4 = immediate + R31
					}
					if count2 == 5 {
						R5 = immediate + R31
					}
					if count2 == 6 {
						R6 = immediate + R31
					}
					if count2 == 7 {
						R7 = immediate + R31
					}
					if count2 == 8 {
						R8 = immediate + R31
					}
					if count2 == 9 {
						R9 = immediate + R31
					}
					if count2 == 10 {
						R10 = immediate + R31
					}
					if count2 == 11 {
						R11 = immediate + R31
					}
					if count2 == 12 {
						R12 = immediate + R31
					}
					if count2 == 13 {
						R13 = immediate + R31
					}
					if count2 == 14 {
						R14 = immediate + R31
					}
					if count2 == 15 {
						R15 = immediate + R31
					}
					if count2 == 16 {
						R16 = immediate + R31
					}
					if count2 == 17 {
						R17 = immediate + R31
					}
					if count2 == 18 {
						R18 = immediate + R31
					}
					if count2 == 19 {
						R19 = immediate + R31
					}
					if count2 == 20 {
						R20 = immediate + R31
					}
					if count2 == 21 {
						R21 = immediate + R31
					}
					if count2 == 22 {
						R22 = immediate + R31
					}
					if count2 == 23 {
						R23 = immediate + R31
					}
					if count2 == 24 {
						R24 = immediate + R31
					}
					if count2 == 25 {
						R25 = immediate + R31
					}
					if count2 == 26 {
						R26 = immediate + R31
					}
					if count2 == 27 {
						R27 = immediate + R31
					}
					if count2 == 28 {
						R28 = immediate + R31
					}
					if count2 == 29 {
						R29 = immediate + R31
					}
					if count2 == 30 {
						R30 = immediate + R31
					}
					if count2 == 31 {
						R31 = immediate + R31
					}
				}
			}
		}

		if len(opcode) == 11 {
			//TEST THIS ONE
			if instructionAddress < 100 {
				instructionString = newLine[40:length]
			} else {
				instructionString = newLine[41:length]
			}

		}

		outputSim += "====================\n"
		outputSim += fmt.Sprintf("cycle:%d \t %d \t %s\n\n", cycle, instructionAddress, instructionString)

		outputSim += "registers:\n"
		outputSim += fmt.Sprintf("r00:\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\n",
			R0, R1, R2, R3, R4, R5, R6, R7)
		outputSim += fmt.Sprintf("r08:\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\n",
			R8, R9, R10, R11, R12, R13, R14, R15)
		outputSim += fmt.Sprintf("r16:\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\n",
			R16, R17, R18, R19, R20, R21, R22, R23)
		outputSim += fmt.Sprintf("r24:\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\n",
			R24, R25, R26, R27, R28, R29, R30, R31)
		outputSim += "data:\n"

		cycle++
		instructionAddress += 4
	}
	//WRITE TO SIMULATOR FILE
	_, err = writeOutputSimFile.WriteString(outputSim)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		// file.Close()
		return
	}
	err = writeOutputFile.Close()
	if err != nil {
		return
	}

}
