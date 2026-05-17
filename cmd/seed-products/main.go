package main

import (
	"log"
	"math/rand"
	"time"

	"ecommerce/internal/database"
	"ecommerce/internal/models"
)

var (
	productNames = []string{
		"Wireless Bluetooth Headphones", "Smart Watch Series X", "Laptop Stand Aluminum",
		"USB-C Hub Multiport", "Mechanical Keyboard RGB", "Wireless Mouse Ergonomic",
		"Portable Power Bank 20000mAh", "Phone Case Premium Leather", "Screen Protector Tempered Glass",
		"Webcam HD 1080p", "External SSD 1TB", "Wireless Charger Fast",
		"Bluetooth Speaker Waterproof", "Gaming Mouse Pad XL", "Monitor Light Bar",
		"Laptop Sleeve 15 inch", "Tablet Stand Adjustable", "Wireless Earbuds Pro",
		"Smart Home Hub", "LED Strip Lights RGB", "Cable Management Kit",
		"Desk Organizer Bamboo", "Laptop Cooling Pad", "Wireless Keyboard Compact",
		"Graphics Tablet Drawing", "Microphone USB Studio", "Ring Light LED",
		"Tripod Smartphone", "Action Camera 4K", "Drone with Camera",
		"Smart TV Box 4K", "Streaming Stick Remote", "Soundbar Wireless",
		"Home Security Camera", "Video Doorbell", "Smart Thermostat",
		"Air Purifier HEPA", "Robot Vacuum Cleaner", "Smart Light Bulb",
		"Coffee Maker Programmable", "Electric Kettle Fast", "Toaster 2-Slice",
		"Blender High Speed", "Food Processor Multi", "Slow Cooker Digital",
		"Air Fryer Large", "Pressure Cooker Electric", "Stand Mixer Professional",
		"Iron Steam", "Garment Steamer", "Vacuum Cleaner Handheld",
		"Floor Lamp Modern", "Ceiling Light LED", "Wall Sconce Set",
		"Rug Area Modern", "Throw Blanket Soft", "Decorative Pillows Set",
		"Wall Art Canvas", "Mirror Round", "Clock Wall Modern",
		"Plant Pot Ceramic", "Artificial Plant Set", "Vase Glass",
		"Curtains Blackout", "Blinds Window", "Shower Curtain Set",
		"Bath Towel Set", "Bath Mat Memory Foam", "Toilet Paper Holder",
		"Kitchen Knife Set", "Cutting Board Bamboo", "Cookware Set Nonstick",
		"Baking Sheet Set", "Mixing Bowls Set", "Measuring Cups Set",
		"Storage Containers Set", "Pantry Organizer", "Spice Rack Wall",
		"Dining Table Set", "Chairs Dining Set", "Bar Stools Set",
		"Sofa Sectional", "Armchair Comfort", "Coffee Table Modern",
		"Bookshelf Tall", "TV Stand Modern", "Console Table Entry",
		"Bed Frame King", "Mattress Memory Foam", "Nightstand Set",
		"Dresser with Mirror", "Wardrobe Closet", "Desk Writing",
		"Office Chair Ergonomic", "Filing Cabinet Lockable", "Bookcase Low",
		"Outdoor Patio Set", "Gazebo Canopy", "Hammock Portable",
		"Grill Gas Propane", "Fire Pit Outdoor", "Lawn Mower Electric",
		"Tool Set Comprehensive", "Drill Cordless", "Ladder Telescoping",
		"Workbench Heavy Duty", "Tool Chest Rolling", "Safety Goggles",
		"Garden Hose Expandable", "Sprinkler System", "Leaf Blower Electric",
		"Hedge Trimmer Cordless", "Chainsaw Electric", "Pressure Washer",
		"Bicycle Mountain", "Skateboard Pro", "Scooter Electric",
		"Helmet Safety", "Knee Pads Set", "Elbow Pads Set",
		"Tent Camping 4-Person", "Sleeping Bag Warm", "Camp Chair Folding",
		"Camping Stove Portable", "Lantern LED", "Cooler Insulated",
		"Fishing Rod Combo", "Tackle Box Large", "Kayak Inflatable",
		"Binoculars HD", "Compass Navigation", "Flashlight Tactical",
		"Yoga Mat Premium", "Dumbbells Set", "Resistance Bands Set",
		"Jump Rope Speed", "Punching Bag Set", "Treadmill Folding",
		"Exercise Bike Stationary", "Rowing Machine", "Elliptical Trainer",
		"Massage Gun Percussion", "Foam Roller High Density", "Balance Board",
		"Sports Water Bottle", "Gym Bag Large", "Towel Microfiber",
		"Running Shoes Pro", "Training Shoes", "Basketball Official",
		"Soccer Ball Pro", "Tennis Racket Pro", "Badminton Set",
		"Swim Goggles Pro", "Snorkel Set", "Surfboard Beginner",
		"Snowboard All-Mountain", "Ski Boots Comfort", "Ice Skates Recreational",
		"Skateboard Complete", "Hoverboard Electric", "Electric Scooter Pro",
		"RC Car High Speed", "Drone Racing FPV", "VR Headset Pro",
		"Gaming Console Latest", "Controller Wireless", "Gaming Headset Surround",
		"Gaming Monitor 144Hz", "Gaming Chair Racing", "Gaming Desk Large",
		"Mechanical Switches Set", "Mouse Bungee", "Headphone Stand",
		"Capture Card 4K", "Streaming Microphone", "Green Screen",
		"Smartphone Case Rugged", "Screen Protector Film", "Wireless Charger Pad",
		"Car Mount Phone", "Bluetooth Adapter Car", "Dash Cam 4K",
		"GPS Navigation", "Radar Detector", "Jump Starter Portable",
		"Car Vacuum Cleaner", "Tire Inflator Digital", "Emergency Kit Roadside",
		"Baby Monitor Video", "Diaper Bag Large", "Baby Carrier Comfort",
		"Crib Convertible", "Changing Table", "Rocking Chair Glider",
		"Stroller Travel System", "Car Seat Infant", "High Chair Modern",
		"Playpen Portable", "Baby Swing Electric", "Bath Tub Baby",
		"Toys Educational Set", "Building Blocks Large", "Puzzle 1000 Pieces",
		"Board Game Collection", "Card Game Premium", "Dice Set Professional",
		"Action Figure Set", "Doll House Deluxe", "Remote Control Robot",
		"Art Supplies Set", "Sketchbook Professional", "Painting Canvas Set",
		"Musical Instrument Set", "Guitar Acoustic", "Keyboard Piano",
		"Drum Set Electronic", "Violin Student", "Microphone Karaoke",
		"Book Collection Classic", "E-Reader Premium", "Audiobook Subscription",
		"Journal Leather", "Pen Set Luxury", "Desk Lamp LED",
		"Backpack Travel", "Luggage Set Spinner", "Suitcase Carry-On",
		"Travel Pillow Memory Foam", "Eye Mask Silk", "Passport Holder",
		"Toiletry Bag Hanging", "Packing Cubes Set", "Travel Adapter Universal",
		"Wallet Leather Slim", "Purse Designer", "Watch Analog Classic",
		"Sunglasses Polarized", "Reading Glasses Blue Light", "Contact Lens Solution",
		"Perfume Luxury", "Cologne Premium", "Skincare Set Complete",
		"Makeup Kit Professional", "Hair Dryer Professional", "Straightener Ceramic",
		"Shaving Kit Premium", "Electric Toothbrush", "Water Flosser",
		"Vitamins Daily", "Supplements Protein", "Fitness Tracker Band",
		"Blood Pressure Monitor", "Thermometer Digital", "Scale Body Fat",
		"First Aid Kit", "Pill Organizer Weekly", "Hot Cold Pack",
		"Air Mattress Queen", "Camping Cot", "Sleeping Pad Insulated",
		"Hammock Double", "Campfire Grill", "Fire Starter Kit",
		"Survival Kit Emergency", "Multi-Tool Swiss", "Paracord 50ft",
		"Compass Lensatic", "Whistle Emergency", "Signal Mirror",
		"Binoculars Compact", "Spotting Scope", "Telescope Astronomy",
		"Microscope Digital", "Telescope Kids", "Science Kit Chemistry",
		"Robotics Kit STEM", "Coding Computer Set", "3D Printer Starter",
		"Drone Camera 4K", "RC Boat Fast", "RC Truck Off-Road",
		"Slot Car Set", "Train Set Electric", "Model Kit Detailed",
		"Painting by Numbers", "Diamond Painting Kit", "Cross Stitch Kit",
		"Knitting Kit Beginner", "Crochet Set Complete", "Sewing Machine Portable",
		"Embroidery Kit", "Quilting Set Starter", "Fabric Bundle Assorted",
		"Jewelry Making Kit", "Beads Set Large", "Wire Wrapping Tools",
		"Leather Crafting Kit", "Wood Carving Set", "Metal Detector",
		"Gold Panning Kit", "Rock Tumbler", "Crystal Growing Kit",
		"Ant Farm Colony", "Butterfly Garden", "Insect Habitat",
		"Bird Feeder Premium", "Bird House Cedar", "Bat House Cedar",
		"Compost Bin Tumbler", "Rain Barrel 50 Gallon", "Garden Tool Set",
		"Greenhouse Mini", "Grow Light LED", "Hydroponic System",
		"Aquarium Kit Complete", "Fish Tank 20 Gallon", "Terrarium Glass",
		"Reptile Habitat", "Small Animal Cage", "Pet Bed Orthopedic",
		"Dog Leash Retractable", "Cat Scratching Post", "Pet Carrier Airline",
		"Automatic Feeder Pet", "Water Fountain Pet", "Grooming Kit Pet",
		"Training Collar Dog", "Litter Box Self-Cleaning", "Aquarium Filter",
		"Fish Food Premium", "Bird Seed Mix", "Small Animal Food",
		"Pet Toy Variety Pack", "Cat Tree Tower", "Dog House Insulated",
		"Horse Blanket Winter", "Saddle Pad English", "Riding Helmet",
		"Equestrian Boots", "Horse Grooming Kit", "Stable Supplies",
		"Fishing Waders", "Hunting Blind", "Tree Stand",
		"Archery Bow Compound", "Crossbow Hunting", "Air Rifle Pellet",
		"Paintball Marker", "Airsoft Gun Electric", "Nerf Gun Elite",
		"Slingshot Professional", "Blowgun Set", "Target Stand",
		"Security Camera Outdoor", "Alarm System Home", "Door Lock Smart",
		"Window Sensor Alarm", "Motion Detector", "Smoke Detector Smart",
		"Carbon Monoxide Detector", "Water Leak Sensor", "Smart Lock Keyless",
		"Safe Fireproof", "Lock Box Portable", "Key Cabinet Wall",
		"Flashlight Rechargeable", "Lantern Camping LED", "Headlamp USB",
		"Emergency Radio Solar", "Power Station Portable", "Generator Inverter",
		"Solar Panel Kit", "Battery Bank Solar", "Charge Controller",
		"Inverter Pure Sine", "Battery Deep Cycle", "Cables Wiring Kit",
		"Fuse Box Replacement", "Circuit Breaker Panel", "Electrical Tape",
		"Wire Strippers Auto", "Multimeter Digital", "Oscilloscope Handheld",
		"Soldering Station", "Heat Gun Professional", "Hot Glue Gun",
		"Drill Press Bench", "Bandsaw Portable", "Table Saw Compact",
		"Miter Saw Sliding", "Circular Saw Cordless", "Jigsaw Orbital",
		"Reciprocating Saw", "Angle Grinder", "Rotary Tool Kit",
		"Belt Sander Bench", "Sander Orbital", "Planer Thickness",
		"Router Table Combo", "Jointer Bench", "Dust Collector",
		"Air Compressor Portable", "Nail Gun Framing", "Staple Gun Heavy",
		"Spray Gun Paint", "Sandblaster Cabinet", "Welder MIG",
		"Plasma Cutter", "Torch Gas", "Welding Helmet Auto",
		"Work Gloves Leather", "Safety Glasses ANSI", "Ear Protection Muffs",
		"Respirator Mask", "Hard Hat Ventilated", "Safety Boots Steel",
		"Tool Belt Heavy Duty", "Apron Welding", "Knee Pads Work",
		"Ladder Extension", "Scaffold Rolling", "Platform Ladder",
		"Pallet Jack Manual", "Hand Truck Folding", "Dolly Heavy Duty",
		"Strapping Kit", "Shrink Wrap System", "Label Maker Industrial",
		"Barcode Scanner", "Receipt Printer", "Cash Register POS",
		"Safe Deposit Box", "Money Counter Machine", "Coin Sorter",
		"Time Clock System", "Badge Printer ID", "Lanyards Bulk",
		"Visitor Management System", "Access Control Keypad", "Intercom System",
		"Video Doorbell Pro", "Smart Lock Hub", "Garage Door Opener",
		"Gate Opener Automatic", "Remote Control Universal", "Smart Plug WiFi",
		"Light Switch Smart", "Dimmer Switch Smart", "Ceiling Fan Smart",
		"Thermostat Programmable", "Air Conditioner Smart", "Heater Space Smart",
		"Humidifier Smart", "Dehumidifier Smart", "Air Purifier Smart",
		"Vacuum Robot Smart", "Mop Robot Smart", "Lawn Mower Robot",
		"Sprinkler Controller Smart", "Pool Cleaner Robot", "Hot Tub Controller",
		"Window Blinds Motorized", "Curtain Motor Smart", "Lock Smart Fingerprint",
		"Camera Security 360", "Doorbell Video 2K", "Light Outdoor Motion",
		"Sensor Door Window", "Siren Alarm Smart", "Keypad Wireless",
		"Hub Smart Home", "Bridge Zigbee", "Range Extender WiFi",
		"Router Gaming", "Modem Cable", "Switch Network Gigabit",
		"Access Point Wireless", "Patch Cable Cat6", "Cable Management Wall",
		"Server Rack 42U", "UPS Battery Backup", "PDU Power Distribution",
		"KVM Switch Dual", "Monitor Mount Dual", "Desk Converter Standing",
		"Ergonomic Footrest", "Laptop Riser Adjustable", "Monitor Arm Single",
		"Webcam 4K Pro", "Microphone Studio USB", "Audio Interface Professional",
		"MIDI Controller Keyboard", "Studio Monitor Speakers", "Headphones Studio",
		"Audio Cable Premium", "Stand Microphone", "Pop Filter",
		"Acoustic Foam Panels", "Studio Light Ring", "Green Screen Muslin",
		"Tripod Camera Heavy", "Camera Lens Prime", "Camera Bag Professional",
		"Memory Card SDXC", "External Hard Drive SSD", "NAS Storage 4-Bay",
		"Cloud Storage Subscription", "Backup Software Pro", "Antivirus Premium",
		"VPN Service Annual", "Password Manager Premium", "Encryption Software",
		"Firewall Hardware", "Network Switch Managed", "Cable Tester Pro",
		"Crimping Tool RJ45", "Punch Down Tool", "Tone Generator",
		"Fiber Optic Kit", "Cable Stripper Coax", "Connectors RF",
		"Satellite Dish", "Antenna TV Digital", "Amplifier Signal",
		"Splitter Cable", "Surge Protector Rack", "Power Strip Smart",
		"Extension Cord Heavy", "Adapter Travel Universal", "Converter Voltage",
		"Battery Rechargeable AA", "Battery Charger Smart", "Power Bank Solar",
		"Generator Portable", "Inverter Power", "Transfer Switch",
		"Fuel Can Portable", "Propane Tank 20lb", "Charcoal Briquettes",
		"Lighter Fluid", "Matches Waterproof", "Fire Starter Magnesium",
		"Lantern Gas", "Stove Camping Propane", "Grill Portable Gas",
		"Cooler Yeti", "Ice Maker Portable", "Freezer Chest",
		"Refrigerator Compact", "Mini Fridge Dorm", "Wine Cooler Built-in",
		"Kegerator Home", "Ice Cream Maker", "Soda Stream",
		"Blender Personal", "Juicer Electric", "Milk Frother",
		"Tea Kettle Electric", "Coffee Grinder", "French Press",
		"Pour Over Set", "Cold Brew Maker", "Espresso Machine",
		"Coffee Maker Drip", "Percolator Stovetop", "Moka Pot",
		"Turkish Coffee Set", "Coffee Warmer", "Travel Mug Insulated",
		"Water Bottle Insulated", "Tumbler Stainless", "Straw Set Silicone",
		"Coaster Set Stone", "Placemats Set", "Table Runner Linen",
		"Napkin Rings Set", "Salt Pepper Mill", "Spice Grinder Electric",
		"Can Opener Electric", "Jar Opener Automatic", "Bottle Opener Wall",
		"Corkscrew Wine", "Wine Aerator", "Decanter Glass",
		"Glassware Set Crystal", "Stemware Rack", "Coaster Holder",
		"Tray Serving Ottoman", "Platter Ceramic", "Bowl Mixing Glass",
		"Measuring Spoon Set", "Kitchen Scale Digital", "Timer Kitchen",
		"Thermometer Candy", "Meat Thermometer Digital", "Oven Thermometer",
		"Pot Holder Silicone", "Oven Mitt Set", "Apron Chef",
		"Dish Drying Rack", "Drain Board Sink", "Soap Dispenser Kitchen",
		"Sponge Holder", "Paper Towel Holder", "Trash Can Kitchen",
		"Recycling Bin Stackable", "Compost Pail Counter", "Dishwasher Safe",
		"Cutting Board Glass", "Cheese Board Set", "Serving Utensils",
		"Salad Tongs Set", "Cake Server Set", "Pie Server",
		"Ice Cream Scoop", "Cookie Scoop", "Melon Baller",
		"Apple Corer", "Cherry Pitter", "Avocado Slicer",
		"Egg Separator", "Egg Poacher Pan", "Pancake Dispenser",
		"Waffle Maker Belgian", "Panini Press", "Sandwich Maker",
		"Toaster Oven Convection", "Air Fryer Toaster", "Rotisserie Oven",
		"Slow Cooker Programmable", "Rice Cooker Digital", "Pressure Cooker Multi",
		"Food Steamer Electric", "Sous Vide Precision", "Dehydrator Food",
		"Vacuum Sealer Food", "Food Slicer Electric", "Meat Grinder",
		"Pasta Maker Manual", "Ravioli Press", "Tortilla Press",
		"Waffle Cone Maker", "Cotton Candy Machine", "Popcorn Machine",
		"Snow Cone Machine", "Slush Machine", "Ice Shaver",
		"Beverage Dispenser", "Drink Mixer", "Blender Commercial",
		"Stand Mixer Commercial", "Food Processor Commercial", "Slicer Commercial",
		"Deep Fryer Commercial", "Griddle Commercial", "Charbroiler Commercial",
		"Hot Plate Commercial", "Steam Table Commercial", "Warmers Commercial",
		"Refrigerator Commercial", "Freezer Commercial", "Ice Machine Commercial",
		"Dishwasher Commercial", "Glass Washer", "Pot Washer",
		"Work Table Stainless", "Shelving Commercial", "Dunnage Rack",
		"Hand Sink Commercial", "Prep Sink", "Three Compartment Sink",
		"Grease Trap", "Hood Exhaust", "Fire Suppression System",
		"Walk-in Cooler", "Walk-in Freezer", "Blast Chiller",
		"Proofing Cabinet", "Holding Cabinet", "Heated Display",
		"Refrigerated Display", "Hot Display Case", "Sushi Display",
		"Bakery Display Case", "Deli Display Case", "Merchandiser Refrigerated",
		"Cash Register Touch", "POS System Restaurant", "Kitchen Display System",
		"Printer Receipt Thermal", "Kitchen Printer", "Label Printer Barcode",
		"Scanner Barcode Wireless", "Scale Digital Portion", "Timer Kitchen Digital",
		"Thermometer Infrared", "Probe Thermometer", "Monitoring System Temp",
		"Camera Kitchen Monitoring", "Music System Restaurant", "Drive-thru Headset",
		"Intercom Drive-thru", "Timer Drive-thru", "Display Menu Digital",
		"Sign LED Programmable", "Poster Frame Lighted", "Banner Stand Retractable",
		"Menu Board Digital", "Table Tent Holder", "Number Stand Table",
		"Coaster Custom Printed", "Napkin Custom Printed", "Matchbook Custom",
		"Toothpick Dispenser", "Straw Dispenser", "Cup Dispenser",
		"Lid Organizer", "Sleeve Cup Dispenser", "Tray Stand Stackable",
		"Bus Tub Round", "Bus Box Rectangular", "Trash Can Commercial",
		"Recycling Bin Commercial", "Mop Bucket Wringer", "Broom Commercial",
		"Dustpan Lobby", "Floor Sign Wet", "Mat Anti-Fatigue",
		"Shoe Covers Disposable", "Gloves Disposable Nitrile", "Hairnets Bouffant",
		"Aprons Disposable", "Beard Net", "First Aid Kit Restaurant",
		"Fire Extinguisher Class K", "Spill Kit Chemical", "Safety Signage",
		"Clock Wall Large", "Calendar Wall Annual", "Whiteboard Wall",
		"Bulletin Board Cork", "Marker Board Magnetic", "Pin Board Fabric",
		"Locker Employee", "Bench Locker Room", "Mirror Full Length",
		"Coat Rack Wall", "Umbrella Stand Floor", "Boot Tray",
		"Floor Mat Entrance", "Rug Runner Hallway", "Door Mat Heavy Duty",
		"Door Stopper Heavy", "Door Knocker Brass", "Mailbox Wall Mount",
		"House Numbers Metal", "Doorbell Chime", "Peephole Door",
		"Lock Deadbolt", "Chain Door Security", "Security Bar Door",
		"Window Lock Sash", "Blind Cord Safety", "Child Safety Lock",
		"Gate Safety Pool", "Alarm Pool", "Cover Pool Safety",
		"Fence Pool Removable", "Ladder Pool Safety", "Buoy Safety Line",
		"Ring Life Preserver", "Throw Bag Rescue", "Hook Life Saving",
		"Whistle Coast Guard", "Flare Signal Kit", "Light Strobe Emergency",
		"Raft Life Inflatable", "Boat Kayak", "Paddle Oar",
		"Motor Boat Electric", "Anchor Boat", "Rope Marine",
		"Chain Anchor", "Bumper Dock", "Fender Boat",
		"Cover Boat", "Trailer Boat", "Winch Boat",
		"Motor Outboard", "Propeller Boat", "Fuel Tank Marine",
		"Bilge Pump", "Livewell Aerator", "Fish Finder GPS",
		"Chartplotter Marine", "VHF Radio Marine", "Radar Marine",
		"AIS Receiver", "Autopilot Marine", "Wind Instrument",
		"Depth Sounder", "Speed Log", "Compass Marine",
		"Binoculars Marine", "Spotting Scope Marine", "Chart Navigation",
		"Book Cruising Guide", "Flag Nautical", "Light Navigation",
		"Horn Boat", "Bell Ship", "Whistle Steam",
	}

	adjectives = []string{
		"Premium", "Deluxe", "Professional", "Ultra", "Advanced", "Smart", "Digital",
		"Wireless", "Portable", "Compact", "Lightweight", "Heavy Duty", "Industrial",
		"Commercial", "Residential", "Outdoor", "Indoor", "Waterproof", "Weatherproof",
		"Shockproof", "Fireproof", "Rustproof", "Scratch Resistant", "Anti-Glare",
		"Anti-Fingerprint", "Anti-Slip", "Ergonomic", "Adjustable", "Foldable", "Collapsible",
		"Retractable", "Extendable", "Telescoping", "Swivel", "Rotating", "Oscillating",
		"Vibrating", "Heated", "Cooled", "Insulated", "Ventilated", "Breathable",
		"Moisture Wicking", "Quick Dry", "Stain Resistant", "Wrinkle Free", "Fade Resistant",
		"UV Protected", "Anti-Bacterial", "Anti-Microbial", "Hypoallergenic", "Organic",
		"Natural", "Synthetic", "Recycled", "Eco-Friendly", "Sustainable", "Biodegradable",
		"Compostable", "Reusable", "Disposable", "Single-Use", "Multi-Use", "Versatile",
		"Universal", "Compatible", "Interchangeable", "Modular", "Customizable", "Personalized",
		"Handmade", "Artisan", "Crafted", "Engineered", "Designed", "Innovative",
		"Revolutionary", "Cutting-Edge", "State-of-the-Art", "Next-Generation", "Future-Ready",
		"Classic", "Vintage", "Retro", "Modern", "Contemporary", "Traditional",
		"Authentic", "Genuine", "Original", "Official", "Licensed", "Certified",
		"Approved", "Tested", "Verified", "Guaranteed", "Warranted", "Insured",
		"Protected", "Secured", "Encrypted", "Private", "Confidential", "Exclusive",
		"Limited Edition", "Special Edition", "Collector's Edition", "Anniversary Edition",
		"Signature Edition", "Pro Edition", "Plus Edition", "Max Edition", "Ultra Edition",
	}

	categories = []string{
		"Electronics", "Computers", "Mobile", "Audio", "Video", "Photography", "Smart Home",
		"Kitchen", "Appliances", "Cookware", "Bakeware", "Dining", "Bar", "Food",
		"Furniture", "Bedroom", "Living Room", "Office", "Outdoor", "Patio", "Garden",
		"Tools", "Hardware", "Automotive", "Sports", "Fitness", "Outdoor Recreation",
		"Camping", "Hiking", "Fishing", "Hunting", "Water Sports", "Winter Sports",
		"Toys", "Games", "Hobbies", "Arts", "Crafts", "Music", "Books",
		"Baby", "Kids", "Pet", "Health", "Wellness", "Beauty", "Fashion",
		"Jewelry", "Watches", "Accessories", "Luggage", "Travel", "Safety", "Security",
	}
)

func generateProductName() string {
	name := productNames[rand.Intn(len(productNames))]
	if rand.Intn(3) == 0 {
		adj := adjectives[rand.Intn(len(adjectives))]
		return adj + " " + name
	}
	return name
}

func generateDescription(name string) string {
	descriptions := []string{
		"High-quality " + name + " designed for everyday use.",
		"Premium " + name + " with advanced features and durability.",
		"Professional-grade " + name + " for demanding applications.",
		"Compact and portable " + name + " perfect for travel.",
		"Versatile " + name + " suitable for multiple purposes.",
		"Ergonomically designed " + name + " for comfort and efficiency.",
		"Innovative " + name + " with cutting-edge technology.",
		"Reliable " + name + " backed by quality craftsmanship.",
		"Stylish " + name + " that combines form and function.",
		"Essential " + name + " for modern living.",
	}
	return descriptions[rand.Intn(len(descriptions))]
}

func generatePrice() float64 {
	return float64(rand.Intn(900)+10) + float64(rand.Intn(99))/100
}

func generateStock() int {
	return rand.Intn(500)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	database.Connect()
	if err := database.DB.AutoMigrate(&models.Product{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Starting to seed 2000 products...")

	// Check existing product count
	var count int64
	database.DB.Model(&models.Product{}).Count(&count)
	targetCount := int64(2000)
	if count >= targetCount {
		log.Printf("Database already has %d products. Target is %d. Skipping seed.", count, targetCount)
		return
	}

	productsNeeded := targetCount - count
	log.Printf("Database has %d products. Adding %d more to reach %d total.", count, productsNeeded, targetCount)

	// Create products in batches
	batchSize := 100
	for i := 0; i < int(productsNeeded); i += batchSize {
		products := make([]models.Product, 0, batchSize)
		for j := 0; j < batchSize && i+j < int(productsNeeded); j++ {
			name := generateProductName()
			products = append(products, models.Product{
				Name:        name,
				Description: generateDescription(name),
				Price:       generatePrice(),
				Stock:       generateStock(),
			})
		}

		if err := database.DB.Create(&products).Error; err != nil {
			log.Fatal("Failed to create products:", err)
		}

		log.Printf("Created %d products (progress: %d/%d)", len(products), count+int64(i)+int64(len(products)), targetCount)
	}

	log.Printf("Successfully seeded products! Total: %d", targetCount)
}
