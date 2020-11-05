using System.Collections.Generic;

namespace Client.Modules.Minimega
{
    public class MinimegaConfiguration : IConfiguration
    {
        // public struct VlanSpecs {         
        public string ManagementVLAN { get => ""; }
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


    }
}