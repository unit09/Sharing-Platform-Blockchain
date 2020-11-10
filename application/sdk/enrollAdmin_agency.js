'use strict';

const FabricCAServices = require('fabric-ca-client');
const { FileSystemWallet, X509WalletMixin } = require('fabric-network');
const fs = require('fs');
const path = require('path');

const ccpPath = path.resolve(__dirname, '..', 'connection-agency.json');
const ccpJSON = fs.readFileSync(ccpPath, 'utf8');
const ccp = JSON.parse(ccpJSON);

console.log(ccp.certificateAuthorities['ca.agency.sharing.com']);

async function main() {

    try {
        const caInfo = ccp.certificateAuthorities['ca.agency.sharing.com'];
        const caTLSCACerts = caInfo.tlsCACerts.pem;
        const ca = new FabricCAServices(caInfo.url, { trustedRoots: caTLSCACerts, verify: false }, caInfo.caName);

        const walletPath = path.join(process.cwd(), '..', 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const adminExists = await wallet.exists('admin_agency');
        if (adminExists) {
            console.log('An identity for the admin user "admin_agency" already exists in the wallet');
            return;
        }

        const enrollment = await ca.enroll({ enrollmentID: 'admin_agency', enrollmentSecret: 'adminpw' });
        const identity = X509WalletMixin.createIdentity('AgencyOrg', enrollment.certificate, enrollment.key.toBytes());
        await wallet.import('admin_agency', identity);
        console.log('Successfully enrolled admin user "admin_agency" and imported it into the wallet');
    } 
    catch (error) {
        console.error(`Failed to enroll admin user "admin_agency": ${error}`);
        process.exit(1);
    }
}

main();