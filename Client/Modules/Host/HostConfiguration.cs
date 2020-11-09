using System.Collections.Generic;
using Hostmodule;

namespace Client.Modules.Host
{
    public class HostConfiguration : IConfiguration<Configuration>
    {
        // public struct VlanSpecs {         
        public string ManagementVLAN { get => "MGMT:101"; }
        public Dictionary<string, string> SnifferVLANs
        {
            get => new Dictionary<string, string>() {
                    {"ADMN","102"},
                    {"MAINT","103"},
                    {"HIL", "104"}
                };
        }
        public Dictionary<string, string> HilVLANs
        {
            get => new Dictionary<string, string>() {
                    {"HIL", "104"}
                };
        }
        public string NetflowTapPort { get => "6661"; }
        public string NetflowTapIP { get => "127.0.0.1"; }
        public string PowerTapPort { get => "6662"; }
        public string SnifferTapName { get => "sniffer_tap"; }
        public string PublisherTapName { get => "publisher_tap"; }
        public string PublisherTapIP { get => "172.31.1.255/16"; }

        public string Location { get => "/path/to/orchestration"; }


        public Configuration ConvertToProtobuf() {

            VlanSpecs vlanSpecs = new VlanSpecs{
                    ManagementVLAN = ManagementVLAN,
            };

            TapSpecs tapSpecs = new TapSpecs{
                    NetflowTapPort = NetflowTapPort,
                    NetflowTapIP = NetflowTapIP,
                    PowerTapPort = PowerTapPort,
                    SnifferTapName = SnifferTapName,
                    PublisherTapName = PublisherTapName,
                    PublisherTapIP = PublisherTapIP
            };

            Orchestrations orchestrations = new Orchestrations{
                Location = Location
            };

            var configuration = new Configuration {
                VlanSpecs = vlanSpecs,
                TapSpecs = tapSpecs,
                Orchestrations = orchestrations
            };

            configuration.VlanSpecs.SnifferVLANs.Add(SnifferVLANs);
            configuration.VlanSpecs.HilVLANs.Add(HilVLANs);

            return configuration;
        }
    }
}