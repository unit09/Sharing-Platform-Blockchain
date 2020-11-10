'use strict';

const { FileSystemWallet, Gateway } = require('fabric-network');
var path = require('path');

const ccpPath_share1 = path.resolve(__dirname, '..', 'connection-share1.json');
const ccpPath_agency = path.resolve(__dirname, '..', 'connection-agency.json');
let ccpPath;
let user;

async function send(type, func, args, res, ccp){
    if(ccp == "share1"){
        ccpPath = ccpPath_share1;
        user = 'user_share1';
    }
    else if(ccp == "agency"){
        ccpPath = ccpPath_agency;
        user = 'user_agency';
    }
    else{
        console.log("unknown org");
        return;
    }

    try {
        const walletPath = path.join(process.cwd(), '..', 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const userExists = await wallet.exists(user);
        if (!userExists) {
            console.log('An identity for the user does not exist in the wallet');
            console.log('Run the registerUser.js application before retrying');
            return;
        }
        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: user, discovery: { enabled: true, asLocalhost: true } });

        const network = await gateway.getNetwork('sharechannel1');
        const contract = network.getContract('share-cc-ch1');

        for(var i = 0; i < args.length; i++){
            args[i] = String(args[i]);
        }

        if(type){
            await contract.submitTransaction(func, ...args);
            console.log('Transaction has been submitted');
            await gateway.disconnect();
            res.send('success');
        }else{
            const result = await contract.evaluateTransaction(func, ...args);
            console.log(`Transaction has been evaluated, result is: ${result.toString()}`);
            res.send(result.toString());
        }
    } catch (error) {
        console.error(`Failed to submit transaction: ${error}`);
        res.send(`Failed to submit transaction: ${error}`);
    }
}

module.exports = {
    send:send
}