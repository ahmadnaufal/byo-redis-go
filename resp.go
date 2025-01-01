package byoredisgo

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	Terminator = "\r\n"

	Echo = "ECHO"
)

const (
	// types: only support bulk string for now
	DataTypeArray      = '*'
	DataTypeBulkString = '$'
)

func handlePayload(data []byte) (BaseDataType, error) {
	res, err := Construct(data)
	if err != nil {
		return nil, err
	}

	// assume that handlePayload will always receive an array
	arr := res.(*Array)
	cmd := arr.Values[0].(*BulkString).Value

	cmds := Command{cmd: cmd, args: arr.Values[1:]}

	return cmds.Handle()
}

type BaseDataType interface {
	Serialize() []byte
	String() string
}

type BulkString struct {
	Value string
}

func (bs *BulkString) Serialize() []byte {
	var bytes []byte
	bytes = append(bytes, DataTypeBulkString)
	bytes = append(bytes, strconv.Itoa(len(bs.Value))...)
	bytes = append(bytes, Terminator...)
	bytes = append(bytes, bs.Value...)
	bytes = append(bytes, Terminator...)

	return bytes
}

func (bs *BulkString) String() string {
	return bs.Value
}

type Array struct {
	Values []BaseDataType
}

func (a *Array) Serialize() []byte {
	var bytes []byte
	bytes = append(bytes, DataTypeArray)
	bytes = append(bytes, strconv.Itoa(len(a.Values))...)
	bytes = append(bytes, Terminator...)

	for _, v := range a.Values {
		b := v.Serialize()
		bytes = append(bytes, b...)
	}

	return bytes
}

func (a *Array) String() string {
	strs := []string{}
	for _, v := range a.Values {
		strs = append(strs, v.String())
	}

	return strings.Join(strs, " ")
}

func Construct(data []byte) (BaseDataType, error) {
	switch data[0] {
	case DataTypeBulkString:
		return parseBulkString(data)
	case DataTypeArray:
		return parseArray(data)
	default:
		return nil, ErrUnrecognizedCommand
	}
}

func parseArray(data []byte) (*Array, error) {
	// temporary workaround: convert to string
	fmt.Println(string(data))
	tokens := strings.SplitN(string(data), Terminator, 1)
	fmt.Println(tokens)

	// get the number of elements
	numElements, err := strconv.Atoi(tokens[0][1:])
	if err != nil {
		return nil, err
	}

	// get the elements
	arr := Array{}
	for i := 1; i <= numElements; i++ {
		// parse the element
		val, err := Construct([]byte(tokens[i]))
		if err != nil {
			return nil, err
		}

		arr.Values = append(arr.Values, val)
	}

	return &arr, nil
}

func parseBulkString(data []byte) (*BulkString, error) {
	// temporary workaround: convert to string
	tokens := strings.Split(string(data), Terminator)

	// get the number of elements
	numElements, err := strconv.Atoi(tokens[0][1:])
	if err != nil {
		return nil, err
	}

	bs := BulkString{Value: tokens[1][:numElements]}

	return &bs, nil
}

type Command struct {
	cmd  string
	args []BaseDataType
}

func (c *Command) Handle() (BaseDataType, error) {
	switch strings.ToUpper(c.cmd) {
	case Echo:
		return c.args[0], nil
	default:
		return nil, ErrUnrecognizedCommand
	}
}
