// Code generated by thriftgo (0.1.7). DO NOT EDIT.

package match_service

import (
	"context"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"strings"
)

type User struct {
	Id    int32  `thrift:"id,1" json:"id"`
	Name  string `thrift:"name,2" json:"name"`
	Score int32  `thrift:"score,3" json:"score"`
}

func NewUser() *User {
	return &User{}
}

func (p *User) GetId() (v int32) {
	return p.Id
}

func (p *User) GetName() (v string) {
	return p.Name
}

func (p *User) GetScore() (v int32) {
	return p.Score
}
func (p *User) SetId(val int32) {
	p.Id = val
}
func (p *User) SetName(val string) {
	p.Name = val
}
func (p *User) SetScore(val int32) {
	p.Score = val
}

var fieldIDToName_User = map[int16]string{
	1: "id",
	2: "name",
	3: "score",
}

func (p *User) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 3:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField3(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_User[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *User) ReadField1(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Id = v
	}
	return nil
}

func (p *User) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Name = v
	}
	return nil
}

func (p *User) ReadField3(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Score = v
	}
	return nil
}

func (p *User) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("User"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}
		if err = p.writeField3(oprot); err != nil {
			fieldId = 3
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *User) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("id", thrift.I32, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Id); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *User) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("name", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Name); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *User) writeField3(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("score", thrift.I32, 3); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteI32(p.Score); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 3 end error: ", p), err)
}

func (p *User) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("User(%+v)", *p)
}

func (p *User) DeepEqual(ano *User) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.Id) {
		return false
	}
	if !p.Field2DeepEqual(ano.Name) {
		return false
	}
	if !p.Field3DeepEqual(ano.Score) {
		return false
	}
	return true
}

func (p *User) Field1DeepEqual(src int32) bool {

	if p.Id != src {
		return false
	}
	return true
}
func (p *User) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Name, src) != 0 {
		return false
	}
	return true
}
func (p *User) Field3DeepEqual(src int32) bool {

	if p.Score != src {
		return false
	}
	return true
}

type Match interface {
	AddUser(ctx context.Context, user *User, info string) (r int32, err error)

	RemoveUser(ctx context.Context, user *User, info string) (r int32, err error)
}

type MatchClient struct {
	c thrift.TClient
}

func NewMatchClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *MatchClient {
	return &MatchClient{
		c: thrift.NewTStandardClient(f.GetProtocol(t), f.GetProtocol(t)),
	}
}

func NewMatchClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *MatchClient {
	return &MatchClient{
		c: thrift.NewTStandardClient(iprot, oprot),
	}
}

func NewMatchClient(c thrift.TClient) *MatchClient {
	return &MatchClient{
		c: c,
	}
}

func (p *MatchClient) Client_() thrift.TClient {
	return p.c
}

func (p *MatchClient) AddUser(ctx context.Context, user *User, info string) (r int32, err error) {
	var _args MatchAddUserArgs
	_args.User = user
	_args.Info = info
	var _result MatchAddUserResult
	if err = p.Client_().Call(ctx, "add_user", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *MatchClient) RemoveUser(ctx context.Context, user *User, info string) (r int32, err error) {
	var _args MatchRemoveUserArgs
	_args.User = user
	_args.Info = info
	var _result MatchRemoveUserResult
	if err = p.Client_().Call(ctx, "remove_user", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

type MatchProcessor struct {
	processorMap map[string]thrift.TProcessorFunction
	handler      Match
}

func (p *MatchProcessor) AddToProcessorMap(key string, processor thrift.TProcessorFunction) {
	p.processorMap[key] = processor
}

func (p *MatchProcessor) GetProcessorFunction(key string) (processor thrift.TProcessorFunction, ok bool) {
	processor, ok = p.processorMap[key]
	return processor, ok
}

func (p *MatchProcessor) ProcessorMap() map[string]thrift.TProcessorFunction {
	return p.processorMap
}

func NewMatchProcessor(handler Match) *MatchProcessor {
	self := &MatchProcessor{handler: handler, processorMap: make(map[string]thrift.TProcessorFunction)}
	self.AddToProcessorMap("add_user", &matchProcessorAddUser{handler: handler})
	self.AddToProcessorMap("remove_user", &matchProcessorRemoveUser{handler: handler})
	return self
}
func (p *MatchProcessor) Process(ctx context.Context, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	name, _, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return false, err
	}
	if processor, ok := p.GetProcessorFunction(name); ok {
		return processor.Process(ctx, seqId, iprot, oprot)
	}
	iprot.Skip(thrift.STRUCT)
	iprot.ReadMessageEnd()
	x := thrift.NewTApplicationException(thrift.UNKNOWN_METHOD, "Unknown function "+name)
	oprot.WriteMessageBegin(name, thrift.EXCEPTION, seqId)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush(ctx)
	return false, x
}

type matchProcessorAddUser struct {
	handler Match
}

func (p *matchProcessorAddUser) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := MatchAddUserArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("add_user", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := MatchAddUserResult{}
	var retval int32
	if retval, err2 = p.handler.AddUser(ctx, args.User, args.Info); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing add_user: "+err2.Error())
		oprot.WriteMessageBegin("add_user", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("add_user", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type matchProcessorRemoveUser struct {
	handler Match
}

func (p *matchProcessorRemoveUser) Process(ctx context.Context, seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := MatchRemoveUserArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("remove_user", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return false, err
	}

	iprot.ReadMessageEnd()
	var err2 error
	result := MatchRemoveUserResult{}
	var retval int32
	if retval, err2 = p.handler.RemoveUser(ctx, args.User, args.Info); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing remove_user: "+err2.Error())
		oprot.WriteMessageBegin("remove_user", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush(ctx)
		return true, err2
	} else {
		result.Success = &retval
	}
	if err2 = oprot.WriteMessageBegin("remove_user", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(ctx); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type MatchAddUserArgs struct {
	User *User  `thrift:"user,1" json:"user"`
	Info string `thrift:"info,2" json:"info"`
}

func NewMatchAddUserArgs() *MatchAddUserArgs {
	return &MatchAddUserArgs{}
}

var MatchAddUserArgs_User_DEFAULT *User

func (p *MatchAddUserArgs) GetUser() (v *User) {
	if !p.IsSetUser() {
		return MatchAddUserArgs_User_DEFAULT
	}
	return p.User
}

func (p *MatchAddUserArgs) GetInfo() (v string) {
	return p.Info
}
func (p *MatchAddUserArgs) SetUser(val *User) {
	p.User = val
}
func (p *MatchAddUserArgs) SetInfo(val string) {
	p.Info = val
}

var fieldIDToName_MatchAddUserArgs = map[int16]string{
	1: "user",
	2: "info",
}

func (p *MatchAddUserArgs) IsSetUser() bool {
	return p.User != nil
}

func (p *MatchAddUserArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_MatchAddUserArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *MatchAddUserArgs) ReadField1(iprot thrift.TProtocol) error {
	p.User = NewUser()
	if err := p.User.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *MatchAddUserArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Info = v
	}
	return nil
}

func (p *MatchAddUserArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("add_user_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *MatchAddUserArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("user", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.User.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *MatchAddUserArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("info", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Info); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *MatchAddUserArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MatchAddUserArgs(%+v)", *p)
}

func (p *MatchAddUserArgs) DeepEqual(ano *MatchAddUserArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.User) {
		return false
	}
	if !p.Field2DeepEqual(ano.Info) {
		return false
	}
	return true
}

func (p *MatchAddUserArgs) Field1DeepEqual(src *User) bool {

	if !p.User.DeepEqual(src) {
		return false
	}
	return true
}
func (p *MatchAddUserArgs) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Info, src) != 0 {
		return false
	}
	return true
}

type MatchAddUserResult struct {
	Success *int32 `thrift:"success,0" json:"success,omitempty"`
}

func NewMatchAddUserResult() *MatchAddUserResult {
	return &MatchAddUserResult{}
}

var MatchAddUserResult_Success_DEFAULT int32

func (p *MatchAddUserResult) GetSuccess() (v int32) {
	if !p.IsSetSuccess() {
		return MatchAddUserResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *MatchAddUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*int32)
}

var fieldIDToName_MatchAddUserResult = map[int16]string{
	0: "success",
}

func (p *MatchAddUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MatchAddUserResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_MatchAddUserResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *MatchAddUserResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Success = &v
	}
	return nil
}

func (p *MatchAddUserResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("add_user_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *MatchAddUserResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI32(*p.Success); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *MatchAddUserResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MatchAddUserResult(%+v)", *p)
}

func (p *MatchAddUserResult) DeepEqual(ano *MatchAddUserResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *MatchAddUserResult) Field0DeepEqual(src *int32) bool {

	if p.Success == src {
		return true
	} else if p.Success == nil || src == nil {
		return false
	}
	if *p.Success != *src {
		return false
	}
	return true
}

type MatchRemoveUserArgs struct {
	User *User  `thrift:"user,1" json:"user"`
	Info string `thrift:"info,2" json:"info"`
}

func NewMatchRemoveUserArgs() *MatchRemoveUserArgs {
	return &MatchRemoveUserArgs{}
}

var MatchRemoveUserArgs_User_DEFAULT *User

func (p *MatchRemoveUserArgs) GetUser() (v *User) {
	if !p.IsSetUser() {
		return MatchRemoveUserArgs_User_DEFAULT
	}
	return p.User
}

func (p *MatchRemoveUserArgs) GetInfo() (v string) {
	return p.Info
}
func (p *MatchRemoveUserArgs) SetUser(val *User) {
	p.User = val
}
func (p *MatchRemoveUserArgs) SetInfo(val string) {
	p.Info = val
}

var fieldIDToName_MatchRemoveUserArgs = map[int16]string{
	1: "user",
	2: "info",
}

func (p *MatchRemoveUserArgs) IsSetUser() bool {
	return p.User != nil
}

func (p *MatchRemoveUserArgs) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 1:
			if fieldTypeId == thrift.STRUCT {
				if err = p.ReadField1(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		case 2:
			if fieldTypeId == thrift.STRING {
				if err = p.ReadField2(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_MatchRemoveUserArgs[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *MatchRemoveUserArgs) ReadField1(iprot thrift.TProtocol) error {
	p.User = NewUser()
	if err := p.User.Read(iprot); err != nil {
		return err
	}
	return nil
}

func (p *MatchRemoveUserArgs) ReadField2(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadString(); err != nil {
		return err
	} else {
		p.Info = v
	}
	return nil
}

func (p *MatchRemoveUserArgs) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("remove_user_args"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField1(oprot); err != nil {
			fieldId = 1
			goto WriteFieldError
		}
		if err = p.writeField2(oprot); err != nil {
			fieldId = 2
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *MatchRemoveUserArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("user", thrift.STRUCT, 1); err != nil {
		goto WriteFieldBeginError
	}
	if err := p.User.Write(oprot); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 1 end error: ", p), err)
}

func (p *MatchRemoveUserArgs) writeField2(oprot thrift.TProtocol) (err error) {
	if err = oprot.WriteFieldBegin("info", thrift.STRING, 2); err != nil {
		goto WriteFieldBeginError
	}
	if err := oprot.WriteString(p.Info); err != nil {
		return err
	}
	if err = oprot.WriteFieldEnd(); err != nil {
		goto WriteFieldEndError
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 2 end error: ", p), err)
}

func (p *MatchRemoveUserArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MatchRemoveUserArgs(%+v)", *p)
}

func (p *MatchRemoveUserArgs) DeepEqual(ano *MatchRemoveUserArgs) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field1DeepEqual(ano.User) {
		return false
	}
	if !p.Field2DeepEqual(ano.Info) {
		return false
	}
	return true
}

func (p *MatchRemoveUserArgs) Field1DeepEqual(src *User) bool {

	if !p.User.DeepEqual(src) {
		return false
	}
	return true
}
func (p *MatchRemoveUserArgs) Field2DeepEqual(src string) bool {

	if strings.Compare(p.Info, src) != 0 {
		return false
	}
	return true
}

type MatchRemoveUserResult struct {
	Success *int32 `thrift:"success,0" json:"success,omitempty"`
}

func NewMatchRemoveUserResult() *MatchRemoveUserResult {
	return &MatchRemoveUserResult{}
}

var MatchRemoveUserResult_Success_DEFAULT int32

func (p *MatchRemoveUserResult) GetSuccess() (v int32) {
	if !p.IsSetSuccess() {
		return MatchRemoveUserResult_Success_DEFAULT
	}
	return *p.Success
}
func (p *MatchRemoveUserResult) SetSuccess(x interface{}) {
	p.Success = x.(*int32)
}

var fieldIDToName_MatchRemoveUserResult = map[int16]string{
	0: "success",
}

func (p *MatchRemoveUserResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *MatchRemoveUserResult) Read(iprot thrift.TProtocol) (err error) {

	var fieldTypeId thrift.TType
	var fieldId int16

	if _, err = iprot.ReadStructBegin(); err != nil {
		goto ReadStructBeginError
	}

	for {
		_, fieldTypeId, fieldId, err = iprot.ReadFieldBegin()
		if err != nil {
			goto ReadFieldBeginError
		}
		if fieldTypeId == thrift.STOP {
			break
		}

		switch fieldId {
		case 0:
			if fieldTypeId == thrift.I32 {
				if err = p.ReadField0(iprot); err != nil {
					goto ReadFieldError
				}
			} else {
				if err = iprot.Skip(fieldTypeId); err != nil {
					goto SkipFieldError
				}
			}
		default:
			if err = iprot.Skip(fieldTypeId); err != nil {
				goto SkipFieldError
			}
		}

		if err = iprot.ReadFieldEnd(); err != nil {
			goto ReadFieldEndError
		}
	}
	if err = iprot.ReadStructEnd(); err != nil {
		goto ReadStructEndError
	}

	return nil
ReadStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read struct begin error: ", p), err)
ReadFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d begin error: ", p, fieldId), err)
ReadFieldError:
	return thrift.PrependError(fmt.Sprintf("%T read field %d '%s' error: ", p, fieldId, fieldIDToName_MatchRemoveUserResult[fieldId]), err)
SkipFieldError:
	return thrift.PrependError(fmt.Sprintf("%T field %d skip type %d error: ", p, fieldId, fieldTypeId), err)

ReadFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T read field end error", p), err)
ReadStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
}

func (p *MatchRemoveUserResult) ReadField0(iprot thrift.TProtocol) error {
	if v, err := iprot.ReadI32(); err != nil {
		return err
	} else {
		p.Success = &v
	}
	return nil
}

func (p *MatchRemoveUserResult) Write(oprot thrift.TProtocol) (err error) {
	var fieldId int16
	if err = oprot.WriteStructBegin("remove_user_result"); err != nil {
		goto WriteStructBeginError
	}
	if p != nil {
		if err = p.writeField0(oprot); err != nil {
			fieldId = 0
			goto WriteFieldError
		}

	}
	if err = oprot.WriteFieldStop(); err != nil {
		goto WriteFieldStopError
	}
	if err = oprot.WriteStructEnd(); err != nil {
		goto WriteStructEndError
	}
	return nil
WriteStructBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
WriteFieldError:
	return thrift.PrependError(fmt.Sprintf("%T write field %d error: ", p, fieldId), err)
WriteFieldStopError:
	return thrift.PrependError(fmt.Sprintf("%T write field stop error: ", p), err)
WriteStructEndError:
	return thrift.PrependError(fmt.Sprintf("%T write struct end error: ", p), err)
}

func (p *MatchRemoveUserResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err = oprot.WriteFieldBegin("success", thrift.I32, 0); err != nil {
			goto WriteFieldBeginError
		}
		if err := oprot.WriteI32(*p.Success); err != nil {
			return err
		}
		if err = oprot.WriteFieldEnd(); err != nil {
			goto WriteFieldEndError
		}
	}
	return nil
WriteFieldBeginError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 begin error: ", p), err)
WriteFieldEndError:
	return thrift.PrependError(fmt.Sprintf("%T write field 0 end error: ", p), err)
}

func (p *MatchRemoveUserResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("MatchRemoveUserResult(%+v)", *p)
}

func (p *MatchRemoveUserResult) DeepEqual(ano *MatchRemoveUserResult) bool {
	if p == ano {
		return true
	} else if p == nil || ano == nil {
		return false
	}
	if !p.Field0DeepEqual(ano.Success) {
		return false
	}
	return true
}

func (p *MatchRemoveUserResult) Field0DeepEqual(src *int32) bool {

	if p.Success == src {
		return true
	} else if p.Success == nil || src == nil {
		return false
	}
	if *p.Success != *src {
		return false
	}
	return true
}
