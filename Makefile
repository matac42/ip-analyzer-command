setup:
	-@multipass stop dev-ip-analyzer-command
	-@multipass delete dev-ip-analyzer-command
	-@multipass purge
	multipass launch --name dev-ip-analyzer-command --mem 4G --disk 10G --cpus 2 --mount .:/home/ubuntu/ip-analyzer 22.04
	multipass exec dev-ip-analyzer-command sudo sh go-setup.sh
	multipass exec dev-ip-analyzer-command sudo sh namespaces/router.sh