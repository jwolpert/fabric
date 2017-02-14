##Hyperledger Jenkins Sandbox Process:

Hyperledger Jenkins Sandbox provides you with a Jenkins Job testing/experimentation environment
that can be used before pushing Job templates to the Production Jenkins UI. It is
configured similarly to the Hyperledger
[ci-management](https://gerrit.hyperledger.org/r/gitweb?p=ci-management.git;a=tree;h=refs/heads/master;hb=refs/heads/master)
master instance, however it cannot publish artifacts or vote in Gerrit. Be aware that
this is a test environment, and as such there a limited allotment of minions to test on before
pushing code to the Hyperledger Fabric repo. Keep the following points in mind
prior to beginning work on Hyperledger Jenkins Sandbox environment:

- Jobs are automatically deleted every weekend
- Committers can login and configure Jenkins jobs in the sandbox directly
- Sandbox jobs CANNOT upload build images to docker hub
- Sandbox jobs CANNOT vote on Gerrit
- Jenkins nodes are configured using Hyperledger openstack infrastructure.

Use this environment to create and execute Jenkins jobs. Before you proceed
further, ensure you have a Linux Foundation ID (LFID), which is required to access Gerrit.
The documentation on requesting a Linux Foundation account is available
[here](http://hyperledger-fabric.readthedocs.io/en/latest/Gerrit/lf-account/).

To download **ci-management**, execute the following command to clone the
**ci-managment** repository.

`git clone ssh://<LFID>@gerrit.hyperledger.org:29418/ci-management && scp -p -P 29418 <LFID>@gerrit.hyperledger.org:hooks/commit-msg ci-management/.git/hooks/`

Once the repository is successfully cloned, the next step is to install JJB (Jenkins Job Builder)
in order to experiment with Jenkins jobs.

###Execute the following commands to install JJB on your machine:

```
cd ci-management
sudo apt-get install python-virtualenv
virtualenv hyp
source hyp/bin/activate
pip install 'jenkins-job-builder==1.5.0'
jenkins-jobs --version
jenkins-jobs test --recursive jjb/
```
### Make a copy of the example JJB config file (in the builder/ directory)

Backup the jenkins.ini.example to jenkins.ini

`cp jenkins.ini.example jenkins.ini`

After copying the jenkins.ini.example, modify `jenkins.ini` with
your **Jenkins LFID username**, **API token** and **Hyperledger jenkins sandbox URL**

```
[job_builder]
ignore_cache=True
keep_descriptions=False
include_path=.:scripts:~/git/
recursive=True

[jenkins]
#user=jenkins
#password=1234567890abcdef1234567890abcdef
#url=http://localhost:8080
user=rameshthoomu <<your LFID username>
password=bbb779809e4669a013b627abca175ed7 <your LFID jenkins sandbox API Token>
url=https://jenkins.hyperledger.org/sandbox
##### This is deprecated, use job_builder section instead
#ignore_cache=True
```
###How to retrieve your API token?
Login to the Jenkins Sandbox environment, go to your user page by clicking on
your username.  Click **Configure** and then click **Show API Token**.

To work on existing jobs or create new jobs, navigate to the `/jjb` directory where
you will find all job templates for the project.  Follow the below commands to test,
update or delete jobs in your sandbox environment.

##To Test a Job:

After you modify or create jobs in the above environment, it's a good practice to
test the job before you update this job in the sandbox environment.

`jenkins-jobs --conf jenkins.ini test jjb/ <job-name>`

**Example:** `jenkins-jobs --conf jenkins.ini test jjb/ fabric-verify-x86_64`

If the job you’d like to test is a template with variables in its name, it must
be manually expanded before use. For example, the commonly used template
`fabric-verify-{arch}` might expand to `fabric-verify-x86-64`.

A successful test will output the XML description of the Jenkins job described by the
specified JJB job name.

Execute the following command to pipeout to a directory:

`jenkins-jobs --conf jenkins.ini test jjb/ <job-name> -o <directoryname>`

The output directory will contain files with the XML configurations.

##To Update a Job:

Ensure you’ve configured your `jenkins.ini` and verified it by outputting valid XML
descriptions of Jenkins jobs. Upon successful verification, execute the following
command to update the job to the Jenkins sandbox.

`jenkins-jobs --conf jenkins.ini update jjb/ <job-name>`

**Example:** `jenkins-jobs --conf jenkins.ini update jjb/ fabric-verify-x86_64`

##Trigger Jobs from Jenkins Sandbox:

Once you push the Jenkins job configuration to the Hyperledger Sandbox environment,
run the job from Jenkins Sandbox webUI. Follow the below process to trigger the build:

Step 1: Login into the [Jenkins Sandbox WebUI](https://jenkins.hyperledger.org/sandbox/)

Step 2: Click on the **Job** which you want to trigger, the click **Build with Parameters**,
and finally click **Build**.

Step 3: Verify the **Build Executor Status** bar and make sure that the build is triggered
on the available executor. In Sandbox you may not see all platforms build
executors and you don't find many like in production CI environment. If you want
to test any of the Hyperledger fabric repositories

Once the Job is triggered, click on the build number to view the job details and the
console output.

## To Delete a Job:

Execute the following command to Delete a job from Sandbox:

`jenkins-jobs --conf jenkins.ini delete jjb/ <job-name>`

**Example:** `jenkins-jobs --conf jenkins.ini delete jjb/ fabric-verify-x86_64`

The above command would delete the `fabric-verify-x86-64` job.

## Modify an Existing Job:

In the Hyperledger Jenkins sandbox, you can directly edit or modify the job
configuration by selecting the Job name and clicking on the **Configure** button.
Then, click the **Apply** and **Save** buttons to save the Job. However, it is
recommended to simply modify the job in your terminal and then follow the
previously described steps in **To Test a Job** and **To Update a Job** to perform
your modifications.
