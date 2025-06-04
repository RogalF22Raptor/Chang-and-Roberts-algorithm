package king
import (
  "encoding/binary"
  "math/rand"
)

type ICandidate interface {
  SelectLeader() int
}

func NewCandidate(id int, input <-chan []byte, output chan<- []byte) ICandidate {
  x := rand.Uint64()
  //println("node", id, x)
  return &Candidate{id,input,output,x}
}

type Candidate struct {
  id int
  input <-chan[]byte
  output chan<- []byte
  randid uint64
}

func (candidate *Candidate) SelectLeader() int {
  
  data:= make([]byte,8*3)
    binary.BigEndian.PutUint64(data[0:8], candidate.randid)
    binary.BigEndian.PutUint64(data[8:16], 0)
    binary.BigEndian.PutUint64(data[16:24], uint64(candidate.id))
    go func(){
      //println("wysylka",candidate.id,candidate.randid,0,candidate.id)
      candidate.output <- data
      
    }()
    
  for {
    datarecive := <- candidate.input
    reciveId:= binary.BigEndian.Uint64(datarecive[0:8])
    ifleader:= binary.BigEndian.Uint64(datarecive[8:16])
    goodid:= binary.BigEndian.Uint64(datarecive[16:24])
    //println("odbiurka", candidate.id,reciveId,ifleader,goodid)
    if(ifleader==2){
      if(candidate.randid == reciveId){
        return int(goodid)
      } else{
        //println("wysylka",candidate.id,reciveId,ifleader,goodid)
        candidate.output <- datarecive
        return int(goodid)
      }
    }
    if(ifleader==1){
      if(candidate.randid == reciveId){
        code:=make([]byte,8*3)
        binary.BigEndian.PutUint64(code[0:8], reciveId)
        binary.BigEndian.PutUint64(code[8:16], 2)
        binary.BigEndian.PutUint64(code[16:24], goodid)
        //println("wysylka",candidate.id,reciveId,2,goodid)
          candidate.output <- code
          
      } else{
        if(candidate.id>int(goodid)){
          code:=make([]byte,8*3)
          binary.BigEndian.PutUint64(code[0:8], reciveId)
          binary.BigEndian.PutUint64(code[8:16], 1)
          binary.BigEndian.PutUint64(code[16:24], uint64(candidate.id))
          //println("wysylka",candidate.id,reciveId,1,candidate.id)
              candidate.output <- code
              
        } else {
          //println("wysylka",candidate.id,reciveId,1,goodid)
            candidate.output <- datarecive
            
        }
      }
    } else if(ifleader==0){
      if(candidate.randid < reciveId){
        code:=make([]byte,8*3)
        binary.BigEndian.PutUint64(code[0:8], reciveId)
        binary.BigEndian.PutUint64(code[8:16], 0)
        binary.BigEndian.PutUint64(code[16:24], uint64(candidate.id))
        //println("wysylka",candidate.id,reciveId,0,candidate.id)
          candidate.output <- code
          
      } else if( candidate.randid == reciveId){
        code:=make([]byte,8*3)
        binary.BigEndian.PutUint64(code[0:8], reciveId)
        binary.BigEndian.PutUint64(code[8:16], 1)
        binary.BigEndian.PutUint64(code[16:24], uint64(candidate.id))
        //println("wysylka kurwa",candidate.id,reciveId,1,uint64(candidate.id))
          candidate.output <- code
          
      }
    }
  }

}
