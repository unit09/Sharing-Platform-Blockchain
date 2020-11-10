'use strict';

const { FileSystemWallet, Gateway, X509WalletMixin } = require('fabric-network');
const path = require('path');

const ccpPath = path.resolve(__dirname, '..', 'connection-agency.json');

async function main() {

    try {
        const walletPath = path.join(process.cwd(), '..', 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const userExists = await wallet.exists('user_agency');
        if (userExists) {
            console.log('An identity for the user "user_agency" already exists in the wallet');
            return;
        }

        const adminExists = await wallet.exists('admin_agency');
        if (!adminExists) {
            console.log('An identity for the admin user "admin_agency" does not exist in the wallet');
            console.log('Run the enrollAdmin.js application before retrying');
            return;
        }

        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'admin_agency', discovery: { enabled: true, asLocalhost: true } });

        const ca = gateway.getClient().getCertificateAuthority();
        const adminIdentity = gateway.getCurrentIdentity();

        const secret = await ca.register({ affiliation: 'org1.department1', enrollmentID: 'user_agency', role: 'client' }, adminIdentity);
        const enrollment = await ca.enroll({ enrollmentID: 'user_agency', enrollmentSecret: secret });
        const userIdentity = X509WalletMixin.createIdentity('AgencyOrg', enrollment.certificate, enrollment.key.toBytes());
        await wallet.import('user_agency', userIdentity);
        console.log('Successfully registered and enrolled admin user "user_agency" and imported it into the wallet');
    } 
    catch (error) {
        console.error(`Failed to register user "user_agency": ${error}`);
        process.exit(1);
    }
}

main();