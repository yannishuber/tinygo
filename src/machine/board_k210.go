// +build maixbit

// Chip datasheet: https://s3.cn-north-1.amazonaws.com.cn/dl.kendryte.com/documents/kendryte_datasheet_20181011163248_en.pdf

package machine

// K210 IO pins.
const (
	P00 Pin = 0
	P01 Pin = 1
	P02 Pin = 2
	P03 Pin = 3
	P04 Pin = 4
	P05 Pin = 5
	P06 Pin = 6
	P07 Pin = 7
	P08 Pin = 8
	P09 Pin = 9
	P10 Pin = 10
	P11 Pin = 11
	P12 Pin = 12
	P13 Pin = 13
	P14 Pin = 14
	P15 Pin = 15
	P16 Pin = 16
	P17 Pin = 17
	P18 Pin = 18
	P19 Pin = 19
	P20 Pin = 20
	P21 Pin = 21
	P22 Pin = 22
	P23 Pin = 23
	P24 Pin = 24
	P25 Pin = 25
	P26 Pin = 26
	P27 Pin = 27
	P28 Pin = 28
	P29 Pin = 29
	P30 Pin = 30
	P31 Pin = 31
	P32 Pin = 32
	P33 Pin = 33
	P34 Pin = 34
	P35 Pin = 35
	P36 Pin = 36
	P37 Pin = 37
	P38 Pin = 38
	P39 Pin = 39
	P40 Pin = 40
	P41 Pin = 41
	P42 Pin = 42
	P43 Pin = 43
	P44 Pin = 44
	P45 Pin = 45
	P46 Pin = 46
	P47 Pin = 47
)

type FPIOAFunc uint16

// FPIOA functions.
const (
	JTAG_TCLK      FPIOAFunc = 0   // JTAG Test Clock
	JTAG_TDI       FPIOAFunc = 1   // JTAG Test Data In
	JTAG_TMS       FPIOAFunc = 2   // JTAG Test Mode Select
	JTAG_TDO       FPIOAFunc = 3   // JTAG Test Data Out
	SPI0_D0        FPIOAFunc = 4   // SPI0 Data 0
	SPI0_D1        FPIOAFunc = 5   // SPI0 Data 1
	SPI0_D2        FPIOAFunc = 6   // SPI0 Data 2
	SPI0_D3        FPIOAFunc = 7   // SPI0 Data 3
	SPI0_D4        FPIOAFunc = 8   // SPI0 Data 4
	SPI0_D5        FPIOAFunc = 9   // SPI0 Data 5
	SPI0_D6        FPIOAFunc = 10  // SPI0 Data 6
	SPI0_D7        FPIOAFunc = 11  // SPI0 Data 7
	SPI0_SS0       FPIOAFunc = 12  // SPI0 Chip Select 0
	SPI0_SS1       FPIOAFunc = 13  // SPI0 Chip Select 1
	SPI0_SS2       FPIOAFunc = 14  // SPI0 Chip Select 2
	SPI0_SS3       FPIOAFunc = 15  // SPI0 Chip Select 3
	SPI0_ARB       FPIOAFunc = 16  // SPI0 Arbitration
	SPI0_SCLK      FPIOAFunc = 17  // SPI0 Serial Clock
	UARTHS_RX      FPIOAFunc = 18  // UART High speed Receiver
	UARTHS_TX      FPIOAFunc = 19  // UART High speed Transmitter
	RESV6          FPIOAFunc = 20  // Reserved function
	RESV7          FPIOAFunc = 21  // Reserved function
	CLK_SPI1       FPIOAFunc = 22  // Clock SPI1
	CLK_I2C1       FPIOAFunc = 23  // Clock I2C1
	GPIOHS0        FPIOAFunc = 24  // GPIO High speed 0
	GPIOHS1        FPIOAFunc = 25  // GPIO High speed 1
	GPIOHS2        FPIOAFunc = 26  // GPIO High speed 2
	GPIOHS3        FPIOAFunc = 27  // GPIO High speed 3
	GPIOHS4        FPIOAFunc = 28  // GPIO High speed 4
	GPIOHS5        FPIOAFunc = 29  // GPIO High speed 5
	GPIOHS6        FPIOAFunc = 30  // GPIO High speed 6
	GPIOHS7        FPIOAFunc = 31  // GPIO High speed 7
	GPIOHS8        FPIOAFunc = 32  // GPIO High speed 8
	GPIOHS9        FPIOAFunc = 33  // GPIO High speed 9
	GPIOHS10       FPIOAFunc = 34  // GPIO High speed 10
	GPIOHS11       FPIOAFunc = 35  // GPIO High speed 11
	GPIOHS12       FPIOAFunc = 36  // GPIO High speed 12
	GPIOHS13       FPIOAFunc = 37  // GPIO High speed 13
	GPIOHS14       FPIOAFunc = 38  // GPIO High speed 14
	GPIOHS15       FPIOAFunc = 39  // GPIO High speed 15
	GPIOHS16       FPIOAFunc = 40  // GPIO High speed 16
	GPIOHS17       FPIOAFunc = 41  // GPIO High speed 17
	GPIOHS18       FPIOAFunc = 42  // GPIO High speed 18
	GPIOHS19       FPIOAFunc = 43  // GPIO High speed 19
	GPIOHS20       FPIOAFunc = 44  // GPIO High speed 20
	GPIOHS21       FPIOAFunc = 45  // GPIO High speed 21
	GPIOHS22       FPIOAFunc = 46  // GPIO High speed 22
	GPIOHS23       FPIOAFunc = 47  // GPIO High speed 23
	GPIOHS24       FPIOAFunc = 48  // GPIO High speed 24
	GPIOHS25       FPIOAFunc = 49  // GPIO High speed 25
	GPIOHS26       FPIOAFunc = 50  // GPIO High speed 26
	GPIOHS27       FPIOAFunc = 51  // GPIO High speed 27
	GPIOHS28       FPIOAFunc = 52  // GPIO High speed 28
	GPIOHS29       FPIOAFunc = 53  // GPIO High speed 29
	GPIOHS30       FPIOAFunc = 54  // GPIO High speed 30
	GPIOHS31       FPIOAFunc = 55  // GPIO High speed 31
	GPIO0          FPIOAFunc = 56  // GPIO pin 0
	GPIO1          FPIOAFunc = 57  // GPIO pin 1
	GPIO2          FPIOAFunc = 58  // GPIO pin 2
	GPIO3          FPIOAFunc = 59  // GPIO pin 3
	GPIO4          FPIOAFunc = 60  // GPIO pin 4
	GPIO5          FPIOAFunc = 61  // GPIO pin 5
	GPIO6          FPIOAFunc = 62  // GPIO pin 6
	GPIO7          FPIOAFunc = 63  // GPIO pin 7
	UART1_RX       FPIOAFunc = 64  // UART1 Receiver
	UART1_TX       FPIOAFunc = 65  // UART1 Transmitter
	UART2_RX       FPIOAFunc = 66  // UART2 Receiver
	UART2_TX       FPIOAFunc = 67  // UART2 Transmitter
	UART3_RX       FPIOAFunc = 68  // UART3 Receiver
	UART3_TX       FPIOAFunc = 69  // UART3 Transmitter
	SPI1_D0        FPIOAFunc = 70  // SPI1 Data 0
	SPI1_D1        FPIOAFunc = 71  // SPI1 Data 1
	SPI1_D2        FPIOAFunc = 72  // SPI1 Data 2
	SPI1_D3        FPIOAFunc = 73  // SPI1 Data 3
	SPI1_D4        FPIOAFunc = 74  // SPI1 Data 4
	SPI1_D5        FPIOAFunc = 75  // SPI1 Data 5
	SPI1_D6        FPIOAFunc = 76  // SPI1 Data 6
	SPI1_D7        FPIOAFunc = 77  // SPI1 Data 7
	SPI1_SS0       FPIOAFunc = 78  // SPI1 Chip Select 0
	SPI1_SS1       FPIOAFunc = 79  // SPI1 Chip Select 1
	SPI1_SS2       FPIOAFunc = 80  // SPI1 Chip Select 2
	SPI1_SS3       FPIOAFunc = 81  // SPI1 Chip Select 3
	SPI1_ARB       FPIOAFunc = 82  // SPI1 Arbitration
	SPI1_SCLK      FPIOAFunc = 83  // SPI1 Serial Clock
	SPI_SLAVE_D0   FPIOAFunc = 84  // SPI Slave Data 0
	SPI_SLAVE_SS   FPIOAFunc = 85  // SPI Slave Select
	SPI_SLAVE_SCLK FPIOAFunc = 86  // SPI Slave Serial Clock
	I2S0_MCLK      FPIOAFunc = 87  // I2S0 Master Clock
	I2S0_SCLK      FPIOAFunc = 88  // I2S0 Serial Clock(BCLK)
	I2S0_WS        FPIOAFunc = 89  // I2S0 Word Select(LRCLK)
	I2S0_IN_D0     FPIOAFunc = 90  // I2S0 Serial Data Input 0
	I2S0_IN_D1     FPIOAFunc = 91  // I2S0 Serial Data Input 1
	I2S0_IN_D2     FPIOAFunc = 92  // I2S0 Serial Data Input 2
	I2S0_IN_D3     FPIOAFunc = 93  // I2S0 Serial Data Input 3
	I2S0_OUT_D0    FPIOAFunc = 94  // I2S0 Serial Data Output 0
	I2S0_OUT_D1    FPIOAFunc = 95  // I2S0 Serial Data Output 1
	I2S0_OUT_D2    FPIOAFunc = 96  // I2S0 Serial Data Output 2
	I2S0_OUT_D3    FPIOAFunc = 97  // I2S0 Serial Data Output 3
	I2S1_MCLK      FPIOAFunc = 98  // I2S1 Master Clock
	I2S1_SCLK      FPIOAFunc = 99  // I2S1 Serial Clock(BCLK)
	I2S1_WS        FPIOAFunc = 100 // I2S1 Word Select(LRCLK)
	I2S1_IN_D0     FPIOAFunc = 101 // I2S1 Serial Data Input 0
	I2S1_IN_D1     FPIOAFunc = 102 // I2S1 Serial Data Input 1
	I2S1_IN_D2     FPIOAFunc = 103 // I2S1 Serial Data Input 2
	I2S1_IN_D3     FPIOAFunc = 104 // I2S1 Serial Data Input 3
	I2S1_OUT_D0    FPIOAFunc = 105 // I2S1 Serial Data Output 0
	I2S1_OUT_D1    FPIOAFunc = 106 // I2S1 Serial Data Output 1
	I2S1_OUT_D2    FPIOAFunc = 107 // I2S1 Serial Data Output 2
	I2S1_OUT_D3    FPIOAFunc = 108 // I2S1 Serial Data Output 3
	I2S2_MCLK      FPIOAFunc = 109 // I2S2 Master Clock
	I2S2_SCLK      FPIOAFunc = 110 // I2S2 Serial Clock(BCLK)
	I2S2_WS        FPIOAFunc = 111 // I2S2 Word Select(LRCLK)
	I2S2_IN_D0     FPIOAFunc = 112 // I2S2 Serial Data Input 0
	I2S2_IN_D1     FPIOAFunc = 113 // I2S2 Serial Data Input 1
	I2S2_IN_D2     FPIOAFunc = 114 // I2S2 Serial Data Input 2
	I2S2_IN_D3     FPIOAFunc = 115 // I2S2 Serial Data Input 3
	I2S2_OUT_D0    FPIOAFunc = 116 // I2S2 Serial Data Output 0
	I2S2_OUT_D1    FPIOAFunc = 117 // I2S2 Serial Data Output 1
	I2S2_OUT_D2    FPIOAFunc = 118 // I2S2 Serial Data Output 2
	I2S2_OUT_D3    FPIOAFunc = 119 // I2S2 Serial Data Output 3
	RESV0          FPIOAFunc = 120 // Reserved function
	RESV1          FPIOAFunc = 121 // Reserved function
	RESV2          FPIOAFunc = 122 // Reserved function
	RESV3          FPIOAFunc = 123 // Reserved function
	RESV4          FPIOAFunc = 124 // Reserved function
	RESV5          FPIOAFunc = 125 // Reserved function
	I2C0_SCLK      FPIOAFunc = 126 // I2C0 Serial Clock
	I2C0_SDA       FPIOAFunc = 127 // I2C0 Serial Data
	I2C1_SCLK      FPIOAFunc = 128 // I2C1 Serial Clock
	I2C1_SDA       FPIOAFunc = 129 // I2C1 Serial Data
	I2C2_SCLK      FPIOAFunc = 130 // I2C2 Serial Clock
	I2C2_SDA       FPIOAFunc = 131 // I2C2 Serial Data
	CMOS_XCLK      FPIOAFunc = 132 // DVP System Clock
	CMOS_RST       FPIOAFunc = 133 // DVP System Reset
	CMOS_PWDN      FPIOAFunc = 134 // DVP Power Down Mode
	CMOS_VSYNC     FPIOAFunc = 135 // DVP Vertical Sync
	CMOS_HREF      FPIOAFunc = 136 // DVP Horizontal Reference output
	CMOS_PCLK      FPIOAFunc = 137 // Pixel Clock
	CMOS_D0        FPIOAFunc = 138 // Data Bit 0
	CMOS_D1        FPIOAFunc = 139 // Data Bit 1
	CMOS_D2        FPIOAFunc = 140 // Data Bit 2
	CMOS_D3        FPIOAFunc = 141 // Data Bit 3
	CMOS_D4        FPIOAFunc = 142 // Data Bit 4
	CMOS_D5        FPIOAFunc = 143 // Data Bit 5
	CMOS_D6        FPIOAFunc = 144 // Data Bit 6
	CMOS_D7        FPIOAFunc = 145 // Data Bit 7
	SCCB_SCLK      FPIOAFunc = 146 // SCCB Serial Clock
	SCCB_SDA       FPIOAFunc = 147 // SCCB Serial Data
	UART1_CTS      FPIOAFunc = 148 // UART1 Clear To Send
	UART1_DSR      FPIOAFunc = 149 // UART1 Data Set Ready
	UART1_DCD      FPIOAFunc = 150 // UART1 Data Carrier Detect
	UART1_RI       FPIOAFunc = 151 // UART1 Ring Indicator
	UART1_SIR_IN   FPIOAFunc = 152 // UART1 Serial Infrared Input
	UART1_DTR      FPIOAFunc = 153 // UART1 Data Terminal Ready
	UART1_RTS      FPIOAFunc = 154 // UART1 Request To Send
	UART1_OUT2     FPIOAFunc = 155 // UART1 User-designated Output 2
	UART1_OUT1     FPIOAFunc = 156 // UART1 User-designated Output 1
	UART1_SIR_OUT  FPIOAFunc = 157 // UART1 Serial Infrared Output
	UART1_BAUD     FPIOAFunc = 158 // UART1 Transmit Clock Output
	UART1_RE       FPIOAFunc = 159 // UART1 Receiver Output Enable
	UART1_DE       FPIOAFunc = 160 // UART1 Driver Output Enable
	UART1_RS485_EN FPIOAFunc = 161 // UART1 RS485 Enable
	UART2_CTS      FPIOAFunc = 162 // UART2 Clear To Send
	UART2_DSR      FPIOAFunc = 163 // UART2 Data Set Ready
	UART2_DCD      FPIOAFunc = 164 // UART2 Data Carrier Detect
	UART2_RI       FPIOAFunc = 165 // UART2 Ring Indicator
	UART2_SIR_IN   FPIOAFunc = 166 // UART2 Serial Infrared Input
	UART2_DTR      FPIOAFunc = 167 // UART2 Data Terminal Ready
	UART2_RTS      FPIOAFunc = 168 // UART2 Request To Send
	UART2_OUT2     FPIOAFunc = 169 // UART2 User-designated Output 2
	UART2_OUT1     FPIOAFunc = 170 // UART2 User-designated Output 1
	UART2_SIR_OUT  FPIOAFunc = 171 // UART2 Serial Infrared Output
	UART2_BAUD     FPIOAFunc = 172 // UART2 Transmit Clock Output
	UART2_RE       FPIOAFunc = 173 // UART2 Receiver Output Enable
	UART2_DE       FPIOAFunc = 174 // UART2 Driver Output Enable
	UART2_RS485_EN FPIOAFunc = 175 // UART2 RS485 Enable
	UART3_CTS      FPIOAFunc = 176 // UART3 Clear To Send
	UART3_DSR      FPIOAFunc = 177 // UART3 Data Set Ready
	UART3_DCD      FPIOAFunc = 178 // UART3 Data Carrier Detect
	UART3_RI       FPIOAFunc = 179 // UART3 Ring Indicator
	UART3_SIR_IN   FPIOAFunc = 180 // UART3 Serial Infrared Input
	UART3_DTR      FPIOAFunc = 181 // UART3 Data Terminal Ready
	UART3_RTS      FPIOAFunc = 182 // UART3 Request To Send
	UART3_OUT2     FPIOAFunc = 183 // UART3 User-designated Output 2
	UART3_OUT1     FPIOAFunc = 184 // UART3 User-designated Output 1
	UART3_SIR_OUT  FPIOAFunc = 185 // UART3 Serial Infrared Output
	UART3_BAUD     FPIOAFunc = 186 // UART3 Transmit Clock Output
	UART3_RE       FPIOAFunc = 187 // UART3 Receiver Output Enable
	UART3_DE       FPIOAFunc = 188 // UART3 Driver Output Enable
	UART3_RS485_EN FPIOAFunc = 189 // UART3 RS485 Enable
	TIMER0_TOGGLE1 FPIOAFunc = 190 // TIMER0 Toggle Output 1
	TIMER0_TOGGLE2 FPIOAFunc = 191 // TIMER0 Toggle Output 2
	TIMER0_TOGGLE3 FPIOAFunc = 192 // TIMER0 Toggle Output 3
	TIMER0_TOGGLE4 FPIOAFunc = 193 // TIMER0 Toggle Output 4
	TIMER1_TOGGLE1 FPIOAFunc = 194 // TIMER1 Toggle Output 1
	TIMER1_TOGGLE2 FPIOAFunc = 195 // TIMER1 Toggle Output 2
	TIMER1_TOGGLE3 FPIOAFunc = 196 // TIMER1 Toggle Output 3
	TIMER1_TOGGLE4 FPIOAFunc = 197 // TIMER1 Toggle Output 4
	TIMER2_TOGGLE1 FPIOAFunc = 198 // TIMER2 Toggle Output 1
	TIMER2_TOGGLE2 FPIOAFunc = 199 // TIMER2 Toggle Output 2
	TIMER2_TOGGLE3 FPIOAFunc = 200 // TIMER2 Toggle Output 3
	TIMER2_TOGGLE4 FPIOAFunc = 201 // TIMER2 Toggle Output 4
	CLK_SPI2       FPIOAFunc = 202 // Clock SPI2
	CLK_I2C2       FPIOAFunc = 203 // Clock I2C2
	INTERNAL0      FPIOAFunc = 204 // Internal function signal 0
	INTERNAL1      FPIOAFunc = 205 // Internal function signal 1
	INTERNAL2      FPIOAFunc = 206 // Internal function signal 2
	INTERNAL3      FPIOAFunc = 207 // Internal function signal 3
	INTERNAL4      FPIOAFunc = 208 // Internal function signal 4
	INTERNAL5      FPIOAFunc = 209 // Internal function signal 5
	INTERNAL6      FPIOAFunc = 210 // Internal function signal 6
	INTERNAL7      FPIOAFunc = 211 // Internal function signal 7
	INTERNAL8      FPIOAFunc = 212 // Internal function signal 8
	INTERNAL9      FPIOAFunc = 213 // Internal function signal 9
	INTERNAL10     FPIOAFunc = 214 // Internal function signal 10
	INTERNAL11     FPIOAFunc = 215 // Internal function signal 11
	INTERNAL12     FPIOAFunc = 216 // Internal function signal 12
	INTERNAL13     FPIOAFunc = 217 // Internal function signal 13
	INTERNAL14     FPIOAFunc = 218 // Internal function signal 14
	INTERNAL15     FPIOAFunc = 219 // Internal function signal 15
	INTERNAL16     FPIOAFunc = 220 // Internal function signal 16
	INTERNAL17     FPIOAFunc = 221 // Internal function signal 17
	CONSTANT       FPIOAFunc = 222 // Constant function
	INTERNAL18     FPIOAFunc = 223 // Internal function signal 18
	DEBUG0         FPIOAFunc = 224 // Debug function 0
	DEBUG1         FPIOAFunc = 225 // Debug function 1
	DEBUG2         FPIOAFunc = 226 // Debug function 2
	DEBUG3         FPIOAFunc = 227 // Debug function 3
	DEBUG4         FPIOAFunc = 228 // Debug function 4
	DEBUG5         FPIOAFunc = 229 // Debug function 5
	DEBUG6         FPIOAFunc = 230 // Debug function 6
	DEBUG7         FPIOAFunc = 231 // Debug function 7
	DEBUG8         FPIOAFunc = 232 // Debug function 8
	DEBUG9         FPIOAFunc = 233 // Debug function 9
	DEBUG10        FPIOAFunc = 234 // Debug function 10
	DEBUG11        FPIOAFunc = 235 // Debug function 11
	DEBUG12        FPIOAFunc = 236 // Debug function 12
	DEBUG13        FPIOAFunc = 237 // Debug function 13
	DEBUG14        FPIOAFunc = 238 // Debug function 14
	DEBUG15        FPIOAFunc = 239 // Debug function 15
	DEBUG16        FPIOAFunc = 240 // Debug function 16
	DEBUG17        FPIOAFunc = 241 // Debug function 17
	DEBUG18        FPIOAFunc = 242 // Debug function 18
	DEBUG19        FPIOAFunc = 243 // Debug function 19
	DEBUG20        FPIOAFunc = 244 // Debug function 20
	DEBUG21        FPIOAFunc = 245 // Debug function 21
	DEBUG22        FPIOAFunc = 246 // Debug function 22
	DEBUG23        FPIOAFunc = 247 // Debug function 23
	DEBUG24        FPIOAFunc = 248 // Debug function 24
	DEBUG25        FPIOAFunc = 249 // Debug function 25
	DEBUG26        FPIOAFunc = 250 // Debug function 26
	DEBUG27        FPIOAFunc = 251 // Debug function 27
	DEBUG28        FPIOAFunc = 252 // Debug function 28
	DEBUG29        FPIOAFunc = 253 // Debug function 29
	DEBUG30        FPIOAFunc = 254 // Debug function 30
	DEBUG31        FPIOAFunc = 255 // Debug function 31
)
