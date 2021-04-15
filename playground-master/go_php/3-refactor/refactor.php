<?php

function findFirstStringInBracket($str)
{
    if (strlen($str) > 0) {
        $pattern = '#\((.*?)\)#';
        echo 'Pattern: ', $pattern, "\n";
        
        preg_match($pattern, $str, $match);
        return $match[1];

    } else {
        return '';
    }
}

$line_test = "(hehe) (haha)";
$result = findFirstStringInBracket($line_test);
assert($result, "hehe");
echo $result, "\n";

$line_test = "hehe) (haha)";
$result = findFirstStringInBracket($line_test);
assert($result, "haha");
echo $result, "\n";

?>