<?php

$file = file($_SERVER['argv'][1]);

$doprint = false;
foreach($file as $line) { 
	if(strpos($line, "== Tick: " . $_SERVER['argv'][2]) === 0) {
		$doprint = true;
		print($line);
	} elseif($doprint) {
		if(strpos($line, '=') === false) {
			print($line);
		} else {
			exit();
		}
	}
}
