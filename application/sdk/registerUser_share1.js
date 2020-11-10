'use strict';

const { FileSystemWallet, Gateway, X509WalletMixin } = require('fabric-network');
const path = require('path');

const ccpPath = path.resolve(__dirname, '..', 'connection-share1.json');

async function main() {

    try {
        const walletPath = path.join(process.cwd(), '..', 'wallet');
        const wallet = new FileSystemWallet(walletPath);
        console.log(`Wallet path: ${walletPath}`);

        const userExists = await wallet.exists('user_share1');
        if (userExists) {
            console.log('An identity for the user "user_share1" already exists in the wallet');
            return;
        }

        const adminExists = await wallet.exists('admin_share1');
        if (!adminExists) {
            console.log('An identity for the admin user "admin_share1" does not exist in the wallet');
            console.log('Run the enrollAdmin.js application before retrying');
            return;
        }

        const gateway = new Gateway();
        await gateway.connect(ccpPath, { wallet, identity: 'admin_share1', discovery: { enabled: true, asLocalhost: true } });

        const ca = gateway.getClient().getCertificateAuthority();
        const adminIdentity = gateway.getCurrentIdentity();

        const secret = await ca.register({ affiliation: 'org1.department1', enrollmentID: 'user_share1', role: 'client' }, adminIdentity);
        const enrollment = await ca.enroll({ enrollmentID: 'user_share1', enrollmentSecret: secret });
        const userIdentity = X509WalletMixin.createIdentity('Share1Org', enrollment.certificate, enrollment.key.toBytes());
        await wallet.import('user_share1', userIdentity);
        console.log('Successfully registered and enrolled admin user "user_share1" and imported it into the wallet');
    } 
    catch (error) {
        console.error(`Failed to register user "user_share1": ${error}`);
        process.exit(1);
    }
}

main();