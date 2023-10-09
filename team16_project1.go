/*******  TO DO    *******/
/* 1-MOVR function with check for bits
   2-Command line execution
```3-2's Complement function/process is incorrect. Needs to check for leading 1 and only *COMPLETED*
	perform 2's complement function if the leading bit is a 1. Otherwise perform normal binary conversion
```4-Formatting is incorrect, on display and write. Should follow example output. See example code in Lecture 6 slides
   5-Break case needs to be added
	- Check that integers are being converted properly according to context (2's C vs Binary convert)
   6-Make use of flags so input/write aren't hardcoded
   7-Test cases need to be generated
   8-Code could be cleaned up and optimized considerably. Several areas like the "write and print" commands
     are redundant and could probably be made more efficient. Struct should probably be used for function variables,
	pointers could reduce the number of copies etc..
	9-First and second source bits are inverted on some functions. Needs to be checked/fixed
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

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
	// println(result)
	return int(result), nil
}

/*********************************************************/
/*********************************************************/
/*******      INSTRUCTION-TYPE FUNCTIONS      ************/
/***************    ALPHABETICALLY    ********************/
/*********************************************************/

/*****************ADD IMMEDIATE FUNCTION*********************/
/*func addImmediate(binaryInstruction string, lineNumber int) {
	instructionType := binaryInstruction[0:11]
	immediateValue := binaryInstruction[11:22]
	registrySource := binaryInstruction[22:27]
	registryDest := binaryInstruction[27:32]
	immediateValueInt, err := twosComplement(immediateValue)
	if error == nil {
		fmt.Println("Error:", err)
	}

}
*/
/*****************ADD FUNCTION*********************/
func addInstruction(binaryInstruction string, lineNumber int) {
	firstSource := binaryInstruction[11:16]
	instructionType := binaryInstruction[0:11]
	valueShamt := binaryInstruction[16:22]
	// Reg One Int Conversion
	firstSourceint, err := binaryToInteger(firstSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	secondSource := binaryInstruction[22:27]
	// Reg Two Int Conversion
	secondSourceint, err := binaryToInteger(secondSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	destinationReg := binaryInstruction[27:32]
	// Reg Three Int Conversion
	destInt, err := binaryToInteger(destinationReg)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%.11s %.5s %.6s %.5s %.5s \t%.1d  ADD R%.1d, R%.1d, R%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, secondSourceint)
	//fmt.Printf("%s %5s %.6s %.5s %.5s \n",
	//firstSource, binaryInstruction, firstSource, firstSource, firstSource)
	/*fmt.Println(binaryInstruction[0:11], "\t", firstSource, "\t", binaryInstruction[16:22], secondSource,
	"\t", destinationReg, lineNumber, " ADD \t", "R", destInt, "R", firstSourceint, "R", secondSourceint)
	*/
	// binaryInstruction = ""    Maybe not needed now

	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s \t%.1d  ADD R%.1d, R%.1d, R%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, secondSourceint)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		// file.Close()
		return
	}

}

/***************ADDI FUNCTION***********************/

func addiInstruction(binaryInstruction string, lineNumber int) {
	firstSource := binaryInstruction[22:27]
	destinationReg := binaryInstruction[27:32]
	instructionType := binaryInstruction[0:10]
	immediateBinary := binaryInstruction[10:22]
	// Reg One Int Conversion
	firstSourceint, err := binaryToInteger(firstSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Reg Two Int Conversion
	immediateInt, err := binaryToInteger(immediateBinary)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Reg Three Int Conversion
	destInt, err := binaryToInteger(destinationReg)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%.10s %.12s %.5s %.5s \t%.1d ADDI R%.1d, R%.1d, #%.1d \n",
		instructionType, immediateBinary, firstSource, destinationReg, lineNumber,
		destInt, firstSourceint, immediateInt)
	//fmt.Println(binaryInstruction[0:11], "\t", firstSource, "\t", binaryInstruction[16:22], secondSource,
	//	"\t", destinationReg, lineNumber, "AND \t", "R", destInt, "R", firstSourceint, "R", secondSourceint)
	// binaryInstruction = ""    Maybe not needed now

	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.10s %.12s %.5s %.5s \t%.1d ADDI R%.1d, R%.1d, #%.1d \n",
		instructionType, immediateBinary, firstSource, destinationReg, lineNumber,
		destInt, firstSourceint, immediateInt)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}
}

/***************AND FUNCTION*******************/
func andInstruction(binaryInstruction string, lineNumber int) {
	firstSource := binaryInstruction[11:16]
	instructionType := binaryInstruction[0:11]
	valueShamt := binaryInstruction[16:22]
	// Reg One Int Conversion
	firstSourceint, err := binaryToInteger(firstSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	secondSource := binaryInstruction[22:27]
	// Reg Two Int Conversion
	secondSourceint, err := binaryToInteger(secondSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	destinationReg := binaryInstruction[27:32]
	// Reg Three Int Conversion
	destInt, err := binaryToInteger(destinationReg)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%.11s %.5s %.6s %.5s %.5s \t%.1d AND R%.1d, R%.1d, R%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, secondSourceint)
	//fmt.Println(binaryInstruction[0:11], "\t", firstSource, "\t", binaryInstruction[16:22], secondSource,
	//	"\t", destinationReg, lineNumber, "AND \t", "R", destInt, "R", firstSourceint, "R", secondSourceint)
	// binaryInstruction = ""    Maybe not needed now
	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s \t%.1d AND R%.1d, R%.1d, R%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, secondSourceint)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}

}

/********************ASR FUNCTION******************************/
func arithShiftFunction(binaryInstruction string, lineNumber int) {
	instructionType := binaryInstruction[0:11]
	firstSource := binaryInstruction[22:27]
	secondSource := binaryInstruction[11:16]
	destinationReg := binaryInstruction[27:32]
	valueShamt := binaryInstruction[16:22]
	// Reg One Int Conversion
	firstSourceint, err := binaryToInteger(firstSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Reg Two Int Conversion
	shamtInt, err := binaryToInteger(valueShamt)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Reg Three Int Conversion
	destInt, err := binaryToInteger(destinationReg)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%.11s %.5s %.6s %.5s %.5s \t%.1d  ASR R%.1d, R%.1d, #%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, shamtInt)
	//fmt.Printf("%s %5s %.6s %.5s %.5s \n",
	//firstSource, binaryInstruction, firstSource, firstSource, firstSource)
	/*fmt.Println(binaryInstruction[0:11], "\t", firstSource, "\t", binaryInstruction[16:22], secondSource,
	"\t", destinationReg, lineNumber, " ADD \t", "R", destInt, "R", firstSourceint, "R", secondSourceint)
	*/
	// binaryInstruction = ""    Maybe not needed now

	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s \t%.1d  ASR R%.1d, R%.1d, #%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, shamtInt)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		// file.Close()
		return
	}

}

/********************BRANCH INSTRUCTION************************/
func branchInstruction(binaryInstruction string, lineNumber int) {
	instructionType := binaryInstruction[0:6]
	bOffset := binaryInstruction[6:32]
	bOffsetInt, err := twosComplement(bOffset)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%.6s %.26s\t%.1d B   #%d \n",
		instructionType, bOffset, lineNumber, bOffsetInt)

	//fmt.Println(instructionType, bOffset, lineNumber, "B", "#", bOffsetInt)
	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.6s %.26s\t%.1d B   #%d \n",
		instructionType, bOffset, lineNumber, bOffsetInt)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}
}

/*******************CONDITIONAL BRANCH NZERO**********************/
func conditionalBranchNz(binaryInstruction string, lineNumber int) {
	instructionType := binaryInstruction[0:8]
	branchOffset := binaryInstruction[8:27]
	bRegistry := binaryInstruction[27:32]
	branchOffsetInt, err := twosComplement(branchOffset)
	bRegistryInt, err := binaryToInteger(bRegistry)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%.8s %.19s %.5s\t%.1d  CBZN R%.1d #%.1d \n",
		instructionType, branchOffset, bRegistry, lineNumber, bRegistryInt, branchOffsetInt) // ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.8s %.19s %.5s\t%.1d CBZN R%.1d #%.1d \n",
		instructionType, branchOffset, bRegistry, lineNumber, bRegistryInt, branchOffsetInt)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}
}

/*******************CONDITIONAL BRANCH ZERO**********************/
func conditionalBranch(binaryInstruction string, lineNumber int) {
	instructionType := binaryInstruction[0:8]
	branchOffset := binaryInstruction[8:27]
	bRegistry := binaryInstruction[27:32]
	bRegistryInt, err := binaryToInteger(bRegistry)
	branchOffsetInt, err := twosComplement(branchOffset)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%.8s %.19s %.5s\t%.1d  CBZ R%.1d #%.1d \n",
		instructionType, branchOffset, bRegistry, lineNumber, bRegistryInt, branchOffsetInt)
	// fmt.Println(instructionType, branchOffset, bRegistry, lineNumber, "CBZ", "R", bRegistryInt, "#", branchOffsetInt)
	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.8s %.19s %.5s\t%.1d CBZ R%.1d #%.1d \n",
		instructionType, branchOffset, bRegistry, lineNumber, bRegistryInt, branchOffsetInt)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}
}

/*****************EOR FUNCTION*********************/
func eorInstruction(binaryInstruction string, lineNumber int) {
	firstSource := binaryInstruction[22:27]
	// Reg One Int Conversion
	firstSourceint, err := binaryToInteger(firstSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	secondSource := binaryInstruction[11:16]
	// Reg Two Int Conversion
	secondSourceint, err := binaryToInteger(secondSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	destinationReg := binaryInstruction[27:32]
	// Reg Three Int Conversion
	destInt, err := binaryToInteger(destinationReg)
	if err != nil {
		fmt.Println("Error:", err)
	}
	instructionType := binaryInstruction[0:11]
	valueShamt := binaryInstruction[16:22]
	fmt.Printf("%.11s %.5s %.6s %.5s %.5s \t%.1d EOR R%.1d, R%.1d, R%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, secondSourceint)
	//fmt.Println(binaryInstruction[0:11], "\t", firstSource, "\t", binaryInstruction[16:22], secondSource,
	//	"\t", destinationReg, lineNumber, "ORR ", " R", destInt, "R", firstSourceint, "R", secondSourceint)
	// binaryInstruction = ""    Maybe not needed now
	// fmt.Printf("%.11s %.5s %.6s %.5s %.5s \t%.1d ORR  R%.1d, R%.1d, R%.1d \n",
	//	instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
	//		destInt, firstSourceint, secondSourceint)
	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s \t%.1d ORR R%.1d, R%.1d, R%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, secondSourceint)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}
}

/******************LDUR FUNCTION*******************/
func ldurInstruction(binaryInstruction string, lineNumber int) {
	offsetValue := binaryInstruction[11:20]
	opCode2 := binaryInstruction[20:22]
	instructionType := binaryInstruction[0:11]
	baseRegistryRn := binaryInstruction[22:27]
	destinationRegRt := binaryInstruction[27:32]
	// Reg One Int Conversion
	baseRegistryRnint, err := binaryToInteger(baseRegistryRn)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Reg Two Int Conversion
	destinationRegRtint, err := binaryToInteger(destinationRegRt)
	if err != nil {
		fmt.Println("Error:", err)
	}
	offsetValueInt, err := binaryToInteger(offsetValue)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%.11s %.9s %.2s %.5s %.5s \t%.1d LDUR R%.1d, [R%.1d, #%.1d] \n",
		instructionType, offsetValue, opCode2, baseRegistryRn, destinationRegRt, lineNumber,
		destinationRegRtint, baseRegistryRnint, offsetValueInt)
	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.11s %.9s %.2s %.5s %.5s \t%.1d  LDUR R%.1d, [R%.1d, #%.1d] \n",
		instructionType, offsetValue, opCode2, baseRegistryRn, destinationRegRt, lineNumber,
		destinationRegRtint, baseRegistryRnint, offsetValueInt)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		// file.Close()
		return
	}

}

/******************LSL FUNCTION********************/
func logicalLeftInstruction(binaryInstruction string, lineNumber int) {
	instructionType := binaryInstruction[0:11]
	firstSource := binaryInstruction[22:27]
	secondSource := binaryInstruction[11:16]
	destinationReg := binaryInstruction[27:32]
	valueShamt := binaryInstruction[16:22]
	// Reg One Int Conversion
	firstSourceint, err := binaryToInteger(firstSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Reg Two Int Conversion
	shamtInt, err := binaryToInteger(valueShamt)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Reg Three Int Conversion
	destInt, err := binaryToInteger(destinationReg)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%.11s %.5s %.6s %.5s %.5s \t%.1d  LSL R%.1d, R%.1d, #%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, shamtInt)
	//fmt.Printf("%s %5s %.6s %.5s %.5s \n",
	//firstSource, binaryInstruction, firstSource, firstSource, firstSource)
	/*fmt.Println(binaryInstruction[0:11], "\t", firstSource, "\t", binaryInstruction[16:22], secondSource,
	"\t", destinationReg, lineNumber, " ADD \t", "R", destInt, "R", firstSourceint, "R", secondSourceint)
	*/
	// binaryInstruction = ""    Maybe not needed now

	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s \t%.1d  LSL R%.1d, R%.1d, #%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, shamtInt)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		// file.Close()
		return
	}

}

/******************LSR FUNCTION********************/
func logicalRightInstruction(binaryInstruction string, lineNumber int) {
	instructionType := binaryInstruction[0:11]
	firstSource := binaryInstruction[22:27]
	secondSource := binaryInstruction[11:16]
	destinationReg := binaryInstruction[27:32]
	valueShamt := binaryInstruction[16:22]
	// Reg One Int Conversion
	firstSourceint, err := binaryToInteger(firstSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Reg Two Int Conversion
	shamtInt, err := binaryToInteger(valueShamt)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Reg Three Int Conversion
	destInt, err := binaryToInteger(destinationReg)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%.11s %.5s %.6s %.5s %.5s \t%.1d  LSR R%.1d, R%.1d, #%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, shamtInt)
	//fmt.Printf("%s %5s %.6s %.5s %.5s \n",
	//firstSource, binaryInstruction, firstSource, firstSource, firstSource)
	/*fmt.Println(binaryInstruction[0:11], "\t", firstSource, "\t", binaryInstruction[16:22], secondSource,
	"\t", destinationReg, lineNumber, " ADD \t", "R", destInt, "R", firstSourceint, "R", secondSourceint)
	*/
	// binaryInstruction = ""    Maybe not needed now

	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s \t%.1d  LSR R%.1d, R%.1d, #%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, shamtInt)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		// file.Close()
		return
	}
}

/*****************MOVK FUNCTION********************/

func movkInstruction(binaryInstruction string, lineNumber int) {
	fieldValue := binaryInstruction[11:27]
	opCode2 := binaryInstruction[9:11]
	instructionType := binaryInstruction[0:9]
	baseRegistry := binaryInstruction[27:32]
	// Reg Int Conversion
	baseRegistryint, err := binaryToInteger(baseRegistry)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Reg Two Int Conversion
	/*opcode2int, err := binaryToInteger(opCode2)
	if err != nil {
		fmt.Println("Error:", err)
	}*/
	//Field Value Conversion
	fieldValueInt, err := binaryToInteger(fieldValue)
	if err != nil {
		fmt.Println("Error:", err)
	}
	bitPattern := 0
	if opCode2 == "00" {
		bitPattern = 0
	} else if opCode2 == "01" {
		bitPattern = 16
	} else if opCode2 == "10" {
		bitPattern = 32
	} else if opCode2 == "11" {
		bitPattern = 48
	}
	fmt.Printf("%.9s %.2s %.16s %.5s \t%.1d MOVK R%.1d, %.1d, LSL %.1d \n",
		instructionType, opCode2, fieldValue, baseRegistry, lineNumber,
		baseRegistryint, fieldValueInt, bitPattern)
	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.9s %.2s %.16s %.5s \t%.1d MOVK R%.1d, %.1d, LSL %.1d] \n",
		instructionType, opCode2, fieldValue, baseRegistry, lineNumber,
		baseRegistryint, fieldValueInt, bitPattern)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		// file.Close()
		return
	}
}

/*****************MOVR FUNCTION********************/

func movrInstruction(binaryInstruction string, lineNumber int) {
	fieldValue := binaryInstruction[11:27]
	opCode2 := binaryInstruction[9:11]
	instructionType := binaryInstruction[0:9]
	baseRegistry := binaryInstruction[27:32]
	// Reg Int Conversion
	baseRegistryint, err := binaryToInteger(baseRegistry)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Reg Two Int Conversion
	//opcode2int, err := binaryToInteger(opCode2)
	if err != nil {
		fmt.Println("Error:", err)
	}
	//Field Value Conversion
	// var fieldValueInt uint32
	fieldValueInt, err := binaryToInteger(fieldValue)
	if err != nil {
		fmt.Println("Error:", err)
	}
	bitPattern := 0
	if opCode2 == "00" {
		bitPattern = 0
	} else if opCode2 == "01" {
		bitPattern = 16
	} else if opCode2 == "10" {
		bitPattern = 32
	} else if opCode2 == "11" {
		bitPattern = 48
	}
	fmt.Printf("%.9s %.2s %.16s %.5s \t%.1d MOVR R%.1d, %.1d, LSL %.1d \n",
		instructionType, opCode2, fieldValue, baseRegistry, lineNumber,
		baseRegistryint, fieldValueInt, bitPattern)
	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.9s %.2s %.16s %.5s \t%.1d MOVK R%.1d, %.1d, LSL %.1d] \n",
		instructionType, opCode2, fieldValue, baseRegistry, lineNumber,
		baseRegistryint, fieldValueInt, bitPattern)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		// file.Close()
		return
	}
}

/*****************ORR FUNCTION*********************/
func orrInstruction(binaryInstruction string, lineNumber int) {
	firstSource := binaryInstruction[11:16]
	// Reg One Int Conversion
	firstSourceint, err := binaryToInteger(firstSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	secondSource := binaryInstruction[22:27]
	// Reg Two Int Conversion
	secondSourceint, err := binaryToInteger(secondSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	destinationReg := binaryInstruction[27:32]
	// Reg Three Int Conversion
	destInt, err := binaryToInteger(destinationReg)
	if err != nil {
		fmt.Println("Error:", err)
	}
	instructionType := binaryInstruction[0:11]
	valueShamt := binaryInstruction[16:22]
	fmt.Printf("%.11s %.5s %.6s %.5s %.5s \t%.1d ORR R%.1d, R%.1d, R%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, secondSourceint)
	//fmt.Println(binaryInstruction[0:11], "\t", firstSource, "\t", binaryInstruction[16:22], secondSource,
	//	"\t", destinationReg, lineNumber, "ORR ", " R", destInt, "R", firstSourceint, "R", secondSourceint)
	// binaryInstruction = ""    Maybe not needed now
	// fmt.Printf("%.11s %.5s %.6s %.5s %.5s \t%.1d ORR  R%.1d, R%.1d, R%.1d \n",
	//	instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
	//		destInt, firstSourceint, secondSourceint)
	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s \t%.1d ORR R%.1d, R%.1d, R%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, secondSourceint)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}
}

/*****************STUR FUNCTION****************/

func sturInstruction(binaryInstruction string, lineNumber int) {
	offsetValue := binaryInstruction[11:20]
	opCode2 := binaryInstruction[20:22]
	instructionType := binaryInstruction[0:11]
	baseRegistryRn := binaryInstruction[22:27]
	destinationRegRt := binaryInstruction[27:32]
	// Reg One Int Conversion
	baseRegistryRnint, err := binaryToInteger(baseRegistryRn)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Reg Two Int Conversion
	destinationRegRtint, err := binaryToInteger(destinationRegRt)
	if err != nil {
		fmt.Println("Error:", err)
	}
	offsetValueInt, err := binaryToInteger(offsetValue)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%.11s %.9s %.2s %.5s %.5s \t%.1d STUR R%.1d, [R%.1d, #%.1d] \n",
		instructionType, offsetValue, opCode2, baseRegistryRn, destinationRegRt, lineNumber,
		destinationRegRtint, baseRegistryRnint, offsetValueInt)
	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.11s %.9s %.2s %.5s %.5s \t%.1d  LDUR R%.1d, [R%.1d, #%.1d] \n",
		instructionType, offsetValue, opCode2, baseRegistryRn, destinationRegRt, lineNumber,
		destinationRegRtint, baseRegistryRnint, offsetValueInt)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		// file.Close()
		return
	}

}

/******************SUB FUNCTION****************/
func subInstruction(binaryInstruction string, lineNumber int) {
	firstSource := binaryInstruction[11:16]
	// Reg One Int Conversion
	firstSourceint, err := binaryToInteger(firstSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	secondSource := binaryInstruction[22:27]
	// Reg Two Int Conversion
	secondSourceint, err := binaryToInteger(secondSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	destinationReg := binaryInstruction[27:32]
	instructionType := binaryInstruction[0:11]
	valueShamt := binaryInstruction[16:22]
	// Reg Three Int Conversion
	destInt, err := binaryToInteger(destinationReg)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%.11s %.5s %.6s %.5s %.5s \t%.1d SUB R%.1d, R%.1d, R%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, secondSourceint)
	//fmt.Println(instructionType, "\t", firstSource, "\t", valueShamt, secondSource,
	//	"\t", destinationReg, lineNumber, "SUB \t", "R", destInt, "R", firstSourceint, "R", secondSourceint)
	// binaryInstruction = ""    Maybe not needed now

	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.11s %.5s %.6s %.5s %.5s \t%.1d SUB R%.1d, R%.1d, R%.1d \n",
		instructionType, firstSource, valueShamt, secondSource, destinationReg, lineNumber,
		destInt, firstSourceint, secondSourceint)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}

}

/******************SUBI FUNCTION********************/

func subiInstruction(binaryInstruction string, lineNumber int) {
	firstSource := binaryInstruction[22:27]
	destinationReg := binaryInstruction[27:32]
	instructionType := binaryInstruction[0:10]
	immediateBinary := binaryInstruction[10:22]
	// Reg One Int Conversion
	firstSourceint, err := binaryToInteger(firstSource)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Reg Two Int Conversion
	immediateInt, err := binaryToInteger(immediateBinary)
	if err != nil {
		fmt.Println("Error:", err)
	}
	// Reg Three Int Conversion
	destInt, err := binaryToInteger(destinationReg)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%.10s %.12s %.5s %.5s \t%.1d SUBI R%.1d, R%.1d, #%.1d \n",
		instructionType, immediateBinary, firstSource, destinationReg, lineNumber,
		destInt, firstSourceint, immediateInt)
	//fmt.Println(binaryInstruction[0:11], "\t", firstSource, "\t", binaryInstruction[16:22], secondSource,
	//	"\t", destinationReg, lineNumber, "AND \t", "R", destInt, "R", firstSourceint, "R", secondSourceint)
	// binaryInstruction = ""    Maybe not needed now
	// ***WRITING TO FILE***
	file, err := os.OpenFile("team16_out_dis.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer file.Close()
	// Write the text to the file
	output := fmt.Sprintf("%.10s %.12s %.5s %.5s \t%.1d SUBI R%.1d, R%.1d, #%.1d \n",
		instructionType, immediateBinary, firstSource, destinationReg, lineNumber,
		destInt, firstSourceint, immediateInt)
	_, err = file.WriteString(output)
	if err != nil {
		fmt.Println("Error writing to the file:", err)
		return
	}
}

/*******************************************************************/
/***********PRIMARY FUNCTION TO READ IN INSTRUCTIONS***************/
/******************************************************************/

func readAndProcessInstructions(binaryInstruction string, lineNumber int) {
	// Check if it matches any known instruction type
	switch binaryInstruction[:6] {
	case "000101":
		branchInstruction(binaryInstruction, lineNumber)
	default:
		switch binaryInstruction[:8] {
		case "10110100":
			conditionalBranch(binaryInstruction, lineNumber)
		case "10110101":
			conditionalBranchNz(binaryInstruction, lineNumber)
		default:
			switch binaryInstruction[:9] {
			case "110100101":
				movrInstruction(binaryInstruction, lineNumber)
			case "111100101":
				movkInstruction(binaryInstruction, lineNumber)
			default:
				switch binaryInstruction[:10] {
				case "1101000100":
					subiInstruction(binaryInstruction, lineNumber)
				case "1001000100":
					addiInstruction(binaryInstruction, lineNumber)
				default:
					switch binaryInstruction[:11] {
					case "10001011000":
						addInstruction(binaryInstruction, lineNumber)
					case "10001010000":
						andInstruction(binaryInstruction, lineNumber)
					case "10101010000":
						orrInstruction(binaryInstruction, lineNumber)
					case "11001011000":
						subInstruction(binaryInstruction, lineNumber)
					case "11111000000":
						sturInstruction(binaryInstruction, lineNumber)
					case "11111000010":
						ldurInstruction(binaryInstruction, lineNumber)
					case "11101010000":
						eorInstruction(binaryInstruction, lineNumber)
					case "11010011100":
						arithShiftFunction(binaryInstruction, lineNumber)
					case "11010011011":
						logicalLeftInstruction(binaryInstruction, lineNumber)
					case "11010011010":
						logicalRightInstruction(binaryInstruction, lineNumber)
					default:
						calc, err := twosComplement(binaryInstruction)
						if err != nil {
							println(calc)
						}
						println(calc)
						//NEED BREAK CASE
					}
				}
			}
		}
	}
}

/************   MAIN   ****************/

func main() {
	var fileInput string
	//user inputs file
	fmt.Print("Enter the input file: ")
	_, err := fmt.Scan(&fileInput)
	if err != nil {
		return
	}
	// Open the input file
	inputFile, err := os.Open(fileInput)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create a reader to read from the input file
	reader := bufio.NewReader(inputFile)
	lineNumber := 96
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
		readAndProcessInstructions(binaryInstruction, lineNumber)
		lineNumber += 4
	}

}
