# mylinewizeuiautomation

Since the mylinewize UI only works when using the Linewize Chrome extension with a managed Chrome profile, the only viable approach I can think of is to use Selenium Grid with Proxmox VMs.

The challenging part is launching a managed Chrome browser that contains the required user profile and extensions. After some research, the only workable solution I found is Selenium Grid. Thanks to my previous work with Kai and David on Linewize agent automation, I gained experience in using Selenium Grid deployed on Proxmox VMs. Both the Go and Node POC implementations represent the minimum working solution.

As for deploying Proxmox VMs to GCP, the main difficulties are creating a customised image and handling potential networking issues (which I haven’t yet reached the stage to address). For now, Proxmox is quite slow, but it remains the only available option.
