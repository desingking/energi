// Copyright 2021 The Energi Core Authors
// This file is part of Energi Core.
//
// Energi Core is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Energi Core is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with Energi Core. If not, see <http://www.gnu.org/licenses/>.

// IMPORTANT: this file is code generated, DO NOT EDIT

package consensus

//go:generate go run ./intervalgen/.

var testDataBlockTimes = []uint64{
  34,55,69,33,66,37,35,81,49,81,86,65,69,84,74,83,47,40,51,31,88,86,57,46,56,53,87,59,82,63,
  48,83,34,66,48,32,89,37,37,54,51,49,69,45,69,63,88,78,46,45,58,50,58,71,53,61,60,37,76,65,
  58,83,38,35,33,37,87,56,52,55,89,45,89,73,74,85,86,58,69,33,77,33,47,32,41,70,55,57,44,41,
  84,31,42,62,69,40,47,48,41,36,37,59,33,86,86,62,77,67,46,36,51,32,64,77,45,69,43,69,31,57,
  51,48,84,69,60,53,42,78,48,83,57,61,83,65,62,42,56,54,53,31,80,76,80,39,59,54,89,83,40,70,
  34,71,87,34,68,75,42,64,32,81,81,56,64,72,87,48,66,70,32,81,34,71,54,48,72,44,43,64,41,70,
  53,37,37,83,77,42,52,56,45,64,66,89,64,30,53,63,30,59,89,32,59,73,63,58,73,63,84,79,60,83,
  81,89,55,37,79,46,69,50,39,89,35,48,40,57,62,78,34,34,66,35,58,54,40,37,56,67,88,53,47,31,
  45,41,50,68,61,44,54,60,65,42,48,48,32,48,62,85,61,57,60,75,30,33,32,36,71,70,47,30,34,64,
  39,64,77,56,54,37,52,77,49,38,79,58,88,87,86,75,52,48,43,67,33,84,42,35,43,42,84,77,36,42,
  69,88,41,59,58,75,73,60,45,36,85,79,35,30,75,78,63,53,38,60,40,39,59,84,78,86,77,75,87,42,
  59,30,78,88,36,53,88,65,86,49,82,89,48,35,67,48,55,66,30,57,60,66,37,68,55,34,51,68,46,64,
  }

var testDataBlockTimeEMA = []uint64{
  59147541,59011556,59339046,58475471,58722177,58009974,57255548,58034054,57737856,58500549,
  59402171,59585706,59894372,60684721,61121288,61838623,61352111,60652042,60335582,59373760,
  60312325,61154544,61018329,60525925,60377534,60135647,61016446,60950333,61640486,61685060,
  61236370,61949932,61033541,61196376,60763708,59820636,60777336,59997751,59243726,59071800,
  58807151,58485605,58830340,58376886,58725185,58865343,59820577,60416624,59943948,59453982,
  59406311,59097908,59061911,59453323,59241738,59299386,59322357,58590476,59161280,59352713,
  59308362,60085137,59361034,58562311,57724203,57044721,58026862,57960407,57764984,57674329,
  58701400,58252173,59260298,59710780,60179279,60993073,61812973,61687958,61927697,60979248,
  61504519,60569945,60125029,59202897,58606081,58979652,58849171,58788542,58303671,57736338,
  58597442,57692608,57178097,57336192,57718612,57137674,56805291,56516593,56007852,55351856,
  54750156,54889495,54171807,55215355,56224688,56414042,57088992,57413943,57039715,56349888,
  56174482,55381876,55664437,56363964,55991375,56417888,55977958,56404911,55571963,55618784,
  55467349,55222518,56166042,56586828,56698736,56577466,56099517,56817566,56528465,57396384,
  57383388,57501966,58337967,58556394,58669299,58122765,58053166,57920275,57758954,56881611,
  57639591,58241572,58954963,58300702,58323630,58181871,59192301,59972881,59318033,59668261,
  58826679,59225804,60136434,59279502,59565420,60071472,59478965,59627195,58721385,59451831,
  60158328,60021989,60152415,60540860,61408373,60968754,61133713,61424410,60459675,61133128,
  60243518,60596189,60379920,59974021,60368315,59831648,59279791,59434552,58830141,59196365,
  58993205,58272116,57574669,58408286,59017851,58459889,58248090,58174382,57742435,57947601,
  58211614,59221069,59377755,58414550,58237023,58393186,57462262,57512680,58545051,57674722,
  57718174,58219217,58375964,58363637,58843518,58979796,59800131,60429634,60415548,61156022,
  61806644,62698229,62445828,61611538,62181651,61651105,61892053,61502150,60764375,61690133,
  60815046,60394881,59726196,59636813,59714294,60313826,59451078,58616617,58858695,58076442,
  58073936,57940364,57352155,56684871,56662416,57001353,58017702,57853187,57497345,56628580,
  56247315,55747403,55558964,55966867,56131888,55734121,55677264,55818993,56120010,55657059,
  55406008,55163188,54403739,54193780,54449721,55451370,55633293,55678103,55819805,56448664,
  55581495,54841118,54092229,53499041,54072843,54595044,54346026,53547796,52906885,53270594,
  52802706,53169830,53951147,54018322,54017721,53459763,53411902,54185283,54015274,53490183,
  54326570,54447010,55547108,56578351,57542996,58115357,57914854,57589777,57111424,57435640,
  56634472,57531703,57022467,56300418,55864339,55409771,56347156,57024299,56334977,55864978,
  56295635,57335122,56799544,56871690,56908684,57501842,58009978,58075225,57646529,56936806,
  57856911,58550127,57777991,56867237,57461754,58135140,58294643,58121048,57461341,57544576,
  56969344,56380185,56466081,57368833,58045265,58961814,59553230,60059682,60942972,60321891,
  60278551,59285812,59899392,60820723,60006928,59777192,60702530,60843431,61668237,61252885,
  61933118,62820556,62334636,61438418,61620765,61174183,60971751,61136612,60115739,60013584,
  60013139,60209430,59448465,59728843,59573799,58735314,58481697,58793773,58374305,58558754,
  }

var testDataBlockTimeDrift = []int64{
  -852459,-988444,-660954,-1524529,-1277823,-1990026,-2744452,-1965946,-2262144,-1499451,
  -597829,-414294,-105628,684721,1121288,1838623,1352111,652042,335582,-626240,
  312325,1154544,1018329,525925,377534,135647,1016446,950333,1640486,1685060,
  1236370,1949932,1033541,1196376,763708,-179364,777336,-2249,-756274,-928200,
  -1192849,-1514395,-1169660,-1623114,-1274815,-1134657,-179423,416624,-56052,-546018,
  -593689,-902092,-938089,-546677,-758262,-700614,-677643,-1409524,-838720,-647287,
  -691638,85137,-638966,-1437689,-2275797,-2955279,-1973138,-2039593,-2235016,-2325671,
  -1298600,-1747827,-739702,-289220,179279,993073,1812973,1687958,1927697,979248,
  1504519,569945,125029,-797103,-1393919,-1020348,-1150829,-1211458,-1696329,-2263662,
  -1402558,-2307392,-2821903,-2663808,-2281388,-2862326,-3194709,-3483407,-3992148,-4648144,
  -5249844,-5110505,-5828193,-4784645,-3775312,-3585958,-2911008,-2586057,-2960285,-3650112,
  -3825518,-4618124,-4335563,-3636036,-4008625,-3582112,-4022042,-3595089,-4428037,-4381216,
  -4532651,-4777482,-3833958,-3413172,-3301264,-3422534,-3900483,-3182434,-3471535,-2603616,
  -2616612,-2498034,-1662033,-1443606,-1330701,-1877235,-1946834,-2079725,-2241046,-3118389,
  -2360409,-1758428,-1045037,-1699298,-1676370,-1818129,-807699,-27119,-681967,-331739,
  -1173321,-774196,136434,-720498,-434580,71472,-521035,-372805,-1278615,-548169,
  158328,21989,152415,540860,1408373,968754,1133713,1424410,459675,1133128,
  243518,596189,379920,-25979,368315,-168352,-720209,-565448,-1169859,-803635,
  -1006795,-1727884,-2425331,-1591714,-982149,-1540111,-1751910,-1825618,-2257565,-2052399,
  -1788386,-778931,-622245,-1585450,-1762977,-1606814,-2537738,-2487320,-1454949,-2325278,
  -2281826,-1780783,-1624036,-1636363,-1156482,-1020204,-199869,429634,415548,1156022,
  1806644,2698229,2445828,1611538,2181651,1651105,1892053,1502150,764375,1690133,
  815046,394881,-273804,-363187,-285706,313826,-548922,-1383383,-1141305,-1923558,
  -1926064,-2059636,-2647845,-3315129,-3337584,-2998647,-1982298,-2146813,-2502655,-3371420,
  -3752685,-4252597,-4441036,-4033133,-3868112,-4265879,-4322736,-4181007,-3879990,-4342941,
  -4593992,-4836812,-5596261,-5806220,-5550279,-4548630,-4366707,-4321897,-4180195,-3551336,
  -4418505,-5158882,-5907771,-6500959,-5927157,-5404956,-5653974,-6452204,-7093115,-6729406,
  -7197294,-6830170,-6048853,-5981678,-5982279,-6540237,-6588098,-5814717,-5984726,-6509817,
  -5673430,-5552990,-4452892,-3421649,-2457004,-1884643,-2085146,-2410223,-2888576,-2564360,
  -3365528,-2468297,-2977533,-3699582,-4135661,-4590229,-3652844,-2975701,-3665023,-4135022,
  -3704365,-2664878,-3200456,-3128310,-3091316,-2498158,-1990022,-1924775,-2353471,-3063194,
  -2143089,-1449873,-2222009,-3132763,-2538246,-1864860,-1705357,-1878952,-2538659,-2455424,
  -3030656,-3619815,-3533919,-2631167,-1954735,-1038186,-446770,59682,942972,321891,
  278551,-714188,-100608,820723,6928,-222808,702530,843431,1668237,1252885,
  1933118,2820556,2334636,1438418,1620765,1174183,971751,1136612,115739,13584,
  13139,209430,-551535,-271157,-426201,-1264686,-1518303,-1206227,-1625695,-1441246,
  }

var testDataBlockTimeIntegral int64 = -602336201
