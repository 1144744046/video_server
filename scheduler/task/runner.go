package task

type Runner struct {
	Controller ControllConn
	Err ControllConn
	Data DataChan
	DataSize int
	LongLived bool
	Dispatcher FN
	Execute FN
}

func NewRunner(size int,longlive bool,d FN,e FN)*Runner{

	return &Runner{
		make(chan string,1),
		make(chan string,1),
		make(chan interface{},size),
		size,
		longlive,
		d,
		e,
	}
}

func (r *Runner)StartDispatch(){
	defer func(){
		if !r.LongLived{
			close(r.Controller)
			close(r.Data)
			close(r.Err)
		}
	}()
	for{
		select{
		case c:=<-r.Controller:
			if c==READY_TO_DISPATCH{
				err:=r.Dispatcher(r.Data)
				if err!=nil{
					r.Err<-CLOSE
				}else{
					r.Controller<-READY_TO_EXECUTE
				}
			}
			if c==READY_TO_EXECUTE{
				err:=r.Execute(r.Data)
				if err!=nil{
					r.Err<-CLOSE
				}else{
					r.Controller<-READY_TO_DISPATCH
				}
			}
		case e:=<-r.Err:
			if e==CLOSE{
				return
			}
		default:

		}
	}
}

func (r *Runner)StartAll(){
	r.Controller<-READY_TO_DISPATCH
	r.StartDispatch()
}