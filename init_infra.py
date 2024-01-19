import logging
from aws_cdk import core
from aws_cdk import aws_eks as eks
from aws_cdk import aws_opensearchservice as opensearch
from aws_cdk import aws_iam as iam

# Setup logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

class MyEksStack(core.Stack):
    def __init__(self, scope: core.Construct, id: str, **kwargs):
        super().__init__(scope, id, **kwargs)

        logger.info("Starting EKS Cluster setup")
        # EKS Cluster Setup
        eks_cluster = eks.Cluster(self, "MyEKSCluster",
            version=eks.KubernetesVersion.V1_21,
            default_capacity=2  # number of instances
        )

        logger.info("EKS Cluster setup completed")

        logger.info("Starting OpenSearch Service setup")
        # OpenSearch Service Setup
        opensearch_domain = opensearch.Domain(self, "MyDomain",
            version=opensearch.EngineVersion.OPENSEARCH_1_0,
            capacity=opensearch.CapacityConfig(
                data_node_instance_type="t3.small.search",
                data_nodes=1
            )
        )

        logger.info("OpenSearch Service setup completed")

        logger.info("Configuring IAM Role and Policy for OpenSearch")
        # IAM Role and Policy for OpenSearch
        role = iam.Role(self, "OpenSearchAccessRole", 
            assumed_by=iam.ServicePrincipal("eks.amazonaws.com")
        )
        policy = iam.PolicyStatement(
            actions=["es:ESHttpGet", "es:ESHttpPut"],
            resources=[opensearch_domain.domain_arn]
        )
        role.add_to_policy(policy)

        logger.info("IAM Role and Policy configuration completed")

        # Post-setup validation or additional configurations can be added here
        # Validate EKS Cluster
        if eks_cluster.cluster_name:
            logger.info(f"EKS Cluster {eks_cluster.cluster_name} created successfully")
        else:
            logger.error("EKS Cluster creation failed")

        # Validate OpenSearch Domain
        if opensearch_domain.domain_name:
            logger.info(f"OpenSearch Domain {opensearch_domain.domain_name} created successfully")
        else:
            logger.error("OpenSearch Domain creation failed")

        # Additional post-setup tasks can be added here

# Initialize a CDK app
app = core.App()

# Instantiate your stack
MyEksStack(app, "MyEksStack")

app.synth()
