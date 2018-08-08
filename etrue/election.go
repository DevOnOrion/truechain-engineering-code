package etrue

import (
	"crypto/ecdsa"
	"github.com/truechain/truechain-engineering-code/core/types"
	"github.com/truechain/truechain-engineering-code/event"
	"github.com/truechain/truechain-engineering-code/core"
	"github.com/truechain/truechain-engineering-code/common"
	"math/big"
	"log"
	"github.com/truechain/truechain-engineering-code/core/fastchain"
	"crypto/sha256"

)

var ( 	z  = 100
 		k  = 10000
)

const (
	fastChainHeadSize = 256
	chainHeadSize  = 4096

)
type VoteuUse struct {
	wi 		int64  //Local value
	seed 	string
	b   	bool
	j 		int

}

type Signature struct{
	blockHash common.Hash
	blockNumber *big.Int
	hash common.Hash
	sign []byte
}

type queue struct{
	queuesize int   //Array size
	head int  		//The header and subscript of the queue
	tail int		//Tail subscript of the queue
	q []int   		//Array
}

type CommitteeMember struct{
	ip		[]string
	port  	uint
	coinbase common.Hash
	pubkey  *ecdsa.PublicKey
}

type Committee struct {
	Id 	*big.Int
	CommitteeList []*Committee
}



type Election struct {
	genesisCommittee []*CommitteeMember

	committee []*CommitteeMember
	xh 			uint	//委员会届数
	fcEvent		fcEvent
	scEvent		scEvent
	flag		bool  //用于处理事件切换的标志位
	number		*big.Int
	fastHead *big.Int
	snailHead *big.Int

	CStartFeed		event.Feed
	CStopFeed		event.Feed
	EStartFeed		event.Feed
	CommitteeFeed	event.Feed
	scope         event.SubscriptionScope

	fastChainHeadCh  chan core.FastChainHeadEvent
	fastChainHeadSub event.Subscription

	chainHeadCh  chan core.ChainHeadEvent
	chainHeadSub event.Subscription

	fastchain *fastchain.FastBlockChain
	snailchain *core.BlockChain
}


//Read creation block information and return public key for signature verification
func  (v VoteuUse)ReadGenesis()[]string{

	return nil
}


//func (e *Election) Events() *event.Feed {
//	return &e.switchFeed
//}

type fcEvent interface {
	AddRemoteFruits([]*types.Block) []error
	PendingFruits() (map[common.Hash]*types.Block, error)
	SubscribeNewFastEvent(chan<- core.FastChainHeadEvent) event.Subscription
}

type scEvent interface {
	AddRemoteFruits([]*types.Block) []error
	PendingFruits() (map[common.Hash]*types.Block, error)
	SubscribeNewChainHeadEvent(chan<- core.ChainHeadEvent) event.Subscription
}

func (e *Election)start(){

	e.fastChainHeadCh = make(chan core.FastChainHeadEvent, fastChainHeadSize )
	e.fastChainHeadSub = e.fcEvent.SubscribeNewFastEvent(e.fastChainHeadCh)


	e.chainHeadCh = make(chan core.ChainHeadEvent, chainHeadSize)
	e.chainHeadSub= e.scEvent.SubscribeNewChainHeadEvent(e.chainHeadCh)

}


//Calculate your own force unit locally
func (v VoteuUse)localForce()int64{


	w := v.wi
	//w_i=(D_pf-〖[h]〗_(-k))/u
	return w

}

//The power function used by the draw function
func powerf(x float64, n int) float64 {
	ans := 1.0

	for n != 0 {
		if n%2 == 1 {
			ans *= x
		}
		x *= x
		n /= 2
	}
	return ans
}

//Factorial function
func factorial(){

}

//The sum function
func sigma(j int,k int,wi int,P int64) {

}

// the draw function is calculated locally for each miner
// the parameters seed, w_i, W, P are required
func sortition()bool{
	//j := 0;
	//
	//for (seed / powerf(2,seedlen)) ^ [Sigma(j,0,wi,P) , Sigma(j+1,0,wi,P)]{
	//
	//j++;
	//
	//if  j > N {
	//return j,true;
	//	}
	//}

	return false;

}

//Used for election counting
func(qu queue) InitQueue(){

	qu.queuesize = z;

	//qu.q = make(int,qu.queuesize)

	qu.tail = 0;

	qu.head = 0;

}

func (qu queue) EnQueue(key int){

	tail := (qu.tail+1) % qu.queuesize //取余保证，当quil=queuesize-1时，再转回0

	if tail == qu.head{

		log.Print("the queue has been filled full!")

	} else {

	qu.q[qu.tail] = key

	qu.tail = tail

	}

}

func (qu queue) gettail() int {

	return qu.tail

}

// Verify checks a raw ECDSA signature.
// Returns true if it's valid and false if not.
func (cm CommitteeMember)Verify(signature []byte)bool {
	// hash message
	digest := sha256.Sum256(signature)
	pubkey := cm.pubkey
	curveOrderByteSize := pubkey.Curve.Params().P.BitLen() / 8

	r, s := new(big.Int), new(big.Int)
	r.SetBytes(signature[:curveOrderByteSize])
	s.SetBytes(signature[curveOrderByteSize:])

	return ecdsa.Verify(pubkey,digest[:], r, s)



}
//Another method for validation
func (cm CommitteeMember)ReVerify(FastHeight *big.Int,FastHash common.Hash, ReceiptHash common.Hash, Sign []byte)bool {
	return true
}
//Verify the fast chain committee signatures in batches
func (cm CommitteeMember) VerifyFastBlockSigns(pvs *[]types.PbftVoteSign) (cfvf *[]types.CommitteeFastSignResult) {
	return cfvf
}

func (e *Election)elect(FastNumber *big.Int, FastHash common.Hash)[]*CommitteeMember {

	return nil
}

func (e *Election)GetCommittee(FastNumber *big.Int, FastHash common.Hash) (*big.Int, []*CommitteeMember){

	//if block, ok := e.blockCache.Get(hash); ok {
	//	return nil,block.(*types.FastBlock)
	//}
	//block := rawdb.ReadBlock(e.db, hash, number)
	//if block == nil {
	//	return nil,nil
	//}
	//
	//e.blockCache.Add(block.Hash(), block)

	return nil, nil
}

func (e *Election)GetXh()uint{
	return e.xh
}

//Monitor both chains and trigger elections at the same time
func (e *Election) loop() {

	// Keep waiting for and reacting to the various events
	//fc := core.BlockChain{}
	Qu := queue{}
	Qu.InitQueue()
	for {
		select {
		// Handle ChainHeadEvent
		case fb := <-e.chainHeadCh:
			if fb.Block != nil {
				//Record Numbers to open elections
				Qu.EnQueue(0)
				if Qu.gettail() == z-1{
						//zl := z-13
						go sortition()

					//bn := fc.GetBlockByNumber(zl)

					e.flag = true
					//e.number = [len(bn)-1].Number()

					e.xh++
				}
			}

		case ev := <-e.fastChainHeadCh:
			if ev.Block != nil{

			}
			if e.flag {

			}
		}
	}
}


func NewElction(fastHead *big.Int,snailHead *big.Int,fastchain *fastchain.FastBlockChain,snailchain *core.BlockChain)*Election {

	e := &Election{
		fastHead:		fastHead,
		snailHead: 		snailHead,
		fastchain:		fastchain,
		snailchain:		snailchain,

	}

	return e
}


//Interface for some subscribed events
type CmmitteeStartEvent struct { start  bool}
type CmmitteeStopEvent struct { stop   bool }
type ElectionStartEvent struct {election bool}
type PbftCommitteeActionEvent struct{ pbftAction *PbftAction}

func (e *Election) SubscribeCmmitteeStartEvent(ch chan<- CmmitteeStartEvent) event.Subscription {
	return e.scope.Track(e.CStopFeed.Subscribe(ch))
}

func (e *Election) SubscribeCmmitteeStopEvent(ch chan<- CmmitteeStopEvent) event.Subscription {
	return e.scope.Track(e.CStopFeed.Subscribe(ch))
}

func (e *Election) SubscribeElectionStartEvent(ch chan<- ElectionStartEvent) event.Subscription {
	return e.scope.Track(e.EStartFeed.Subscribe(ch))
}

func (e *Election) SubscribeCommitteeActionEvent(ch chan<- PbftCommitteeActionEvent) event.Subscription {
	return e.scope.Track(e.CommitteeFeed.Subscribe(ch))
}