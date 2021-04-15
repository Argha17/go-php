<?php

function findFirstStringInBracket($str)
{
    if (strlen($str) > 0) {
        $firstbracket = strstr($str, '(');
        if ($firstbracket) {
            $firstbracket = ltrim($firstbracket, '(');
            return strstr($firstbracket, ')', true);
        } else {
            return '';
        }
    } else {
        return '';
    }
}

$line_test = "(hehe) (haha)";
echo findFirstStringInBracket($line_test) ,"\n";

?>