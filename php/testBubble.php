<?php
$init = round(microtime(true) * 1000);

for ($c = 0; $c < 1000000; $c++) {
	bubble();
}

$end= round(microtime(true) * 1000);
echo ($end - $init);

function bubble() {
	$arrayb = [ 
			3,
			4,
			1,
			3,
			5,
			1,
			92,
			2,
			4124,
			424,
			52,
			12 
	];
	$size = sizeof($arrayb);
// 	for($i = 0; $i < $size; $i ++) {
// 		for($y = 0; $y < ($size - 1); $y ++) {
// 			if ($arrayb [$y + 1] < $arrayb [$y]) {
// 				$t = $arrayb [$y];
// 				$arrayb [$y] = $arrayb [$y + 1];
// 				$arrayb [$y + 1] = $t;
// 			}
// 		}
// 	}
	sort($arrayb);
}

