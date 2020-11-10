'use strict';

const FabricCAServices = require('fabric-ca-client');
const { FileSystemWallet, X509WalletMixin } = require('fabric-network');
const fs = require('fs');
const path = require('path');

const ccpPath = path.resolve(__dirname, '..', 'connection-share1.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

console.log(ccp.certificateAuthorities['ca.share1.sharing.com']);

async function main() {

    try {
        const caInfo = ccp.certificateAuthorities['ca.share1.sharing.com'];
        const caTLSCACerts = caInfo.tlsCACerts.pem;
        const ca = new FabricCAServices(caInfo.url, { trustedRoots: caTLSCACerts, verify: false }, caInfo.caName);

        const walletPath = path.join(process.cwd(), '..', 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const adminExists = await wallet.exists('admin_share1');
        if (adminExists) {
            console.log('An identity for the admin user "admin_share1" already exists in the wallet');
            return;
        }

        const enrollment = await ca.enroll({ enrollmentID: 'admin_share1', enrollmentSecret: 'adminpw' });
        const identity = X509WalletMixin.createIdentity('Share1Org', enrollment.certificate, enrollment.key.toBytes());
        await wallet.import('admin_share1', identity);
        console.log('Successfully enrolled admin user "admin_share1" and imported it into the wallet');
    } 
    catch (error) {
        console.error(`Failed to enroll admin user "admin_share1": ${error}`);
        process.exit(1);
    }
}

main();