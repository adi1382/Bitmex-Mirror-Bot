updateRemoveButtonStatus()

function addSubAccount() {
    let nSubs = getNumberOfSubAccounts();
    let newId = "subAccount"+(nSubs+1).toString()
    // console.log(newId)
    // alert(newId)
    let newAccount = document.getElementById("subAccount1").cloneNode(true)
    newAccount.id = newId

    // console.log(newAccount.childNodes[0])

    document.getElementById("subAccounts").appendChild(newAccount)
    document.getElementById(newId).getElementsByTagName("h6")[0].textContent = "Account "+(nSubs+1).toString()

    document.getElementById(newId).getElementsByTagName("select")[0].name = "sub"+(nSubs+1).toString()+"_Status"
    document.getElementById(newId).getElementsByTagName("select")[1].name = "sub"+(nSubs+1).toString()+"_CopyLeverage"
    document.getElementById(newId).getElementsByTagName("select")[2].name = "sub"+(nSubs+1).toString()+"_BalanceProportional"


    document.getElementById(newId).getElementsByTagName("input")[0].name = "sub"+(nSubs+1).toString()+"_FixedProportion"
    document.getElementById(newId).getElementsByTagName("input")[1].name = "sub"+(nSubs+1).toString()+"_ApiKey"
    document.getElementById(newId).getElementsByTagName("input")[2].name = "sub"+(nSubs+1).toString()+"_Secret"
    document.getElementById(newId).getElementsByTagName("input")[3].name = "sub"+(nSubs+1).toString()+"_AccountName"

    let newElements = document.getElementById(newId).getElementsByClassName("form-control")
    newElements[0].value = 0
    newElements[1].value = 1
    newElements[2].value = 0
    newElements[3].value = 1
    newElements[4].value = ""
    newElements[5].value = ""
    newElements[6].value = ""

    updateRemoveButtonStatus()
    console.log("Completed")
}

function removeSubAccount() {
    let nSubs = getNumberOfSubAccounts()

    if (nSubs>1) {
        let lastId = "subAccount"+nSubs.toString()
        document.getElementById(lastId).remove()
    }

    updateRemoveButtonStatus()
}

function getNumberOfSubAccounts() {
    let counter = 0;
    for (let i=1; i<Infinity; i++) {
        if (document.getElementById("subAccount"+i.toString()) != null) {
            counter++;
        } else {
            break;
        }
    }
    return counter;
}

function updateRemoveButtonStatus() {
    document.getElementById("removeSubAccountButton").disabled = getNumberOfSubAccounts() <= 1;
}