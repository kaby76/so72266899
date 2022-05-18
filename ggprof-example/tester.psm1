# Template generated code from trgen 0.16.4
function Build-Grammar {
    $g = antlr parser/SimpleBoolean.g4 -encoding utf-8 -Dlanguage=Go -o parser -lib parser -package parser 
    if($LASTEXITCODE -ne 0){
        return @{
            Message = $g
            Success = $false
        }
    }
    $g = antlr parser/SimpleBoolean.g4 -encoding utf-8 -Dlanguage=Go -o parser -lib parser -package parser 
    if($LASTEXITCODE -ne 0){
        return @{
            Message = $g
            Success = $false
        }
    }

    # Output version of pwsh.
    #Get-Host | Select-Object Version | Write-Host
    # Output environmental variables.
    #dir env: | Out-String | Write-Host
    # Output go version
    #go version | Write-Host
    $env:GO111MODULE = "on"
    $g = go get github.com/antlr/antlr4/runtime/Go/antlr@4.10
    if($LASTEXITCODE -ne 0){
        return @{
            Message = $g
            Success = $false
        }
    }
    $msg = go build Test.go
    return @{
        Message = $msg
        Success = $LASTEXITCODE -eq 0
    }
}

function Test-Case {
    param (
        $InputFile,
        $TokenFile,
        $TreeFile,
        $ErrorFile
    )
    $o = trwdog ./Test -file $InputFile
    $failed = $LASTEXITCODE -ne 0
    if ($failed -and $errorFile) {
        return $true
    }
    if(!$failed -and !$errorFile){
        return $true
    }
    return $false
}