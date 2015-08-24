# Go Development Env

This repo helps you get started with Go development with a Vagrant env.

Requirements
 - vagrant : Latest
 - ansible : Must be installed on host for provisioning the vm


The VM gets a the IP defined in the Vagrant file assigned.

  ```

    config.vm.network :private_network, ip: "192.168.19.20"

  ```

The sync directory with the VM is located at,

 

        config.vm.synced_folder ".", "/home/vagrant/go/project"

 

- The provisioning scripts are still immature, but you can hack and fix any issues if you are brave. Don't forget to send a pull request if you do.
